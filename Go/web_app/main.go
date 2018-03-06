package main

import (
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/codegangsta/negroni"
	gmux "github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/gorp.v1"
)

type Book struct {
	// The `db:"xxx" tells gorp the structure of our database, also the primary key row MUST be int64`
	PK             int64  `db:"pk"`
	Title          string `db:"title"`
	Author         string `db:"author"`
	Classification string `db:"classification"`
	ID             string `db:"id"`
}

type Page struct {
	Books []Book
}

// These 'struct tags' (the weird `` xml thing) tells the decorder how to populate this struct from its
// XML counterpart
type SearchResult struct {
	Title  string `xml:"title,attr"`
	Author string `xml:"author,attr"`
	Year   string `xml:"hyr,attr"`
	ID     string `xml:"owi,attr"`
}

type ClassifyBookResponse struct {
	BookData struct {
		Title  string `xml:"title,attr"`
		Author string `xml:"author,attr"`
		ID     string `xml:"owi,attr"`
	} `xml:"work"`
	Classification struct {
		MostPopular string `xml:"sfa,attr"`
	} `xml:"recommendations>ddc>mostPopular"`
}

var db *sql.DB
var dbmap *gorp.DbMap

func initDb() {
	db, _ = sql.Open("sqlite3", "dev.db")
	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	// Tell gorp we expect a table that matches our book object, called "books"
	//.SetKeys tells gorp that our primary key is auto incremented and is called "pk"
	dbmap.AddTableWithName(Book{}, "books").SetKeys(true, "pk")
	// This will create the table if it doesn't already exist
	dbmap.CreateTablesIfNotExists()
}

// next http.HandlerFunc - Notice it's "Handler", not "Handle"
func verifyDataBase(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	err := db.Ping()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		// WE didn't use 'return' in our previous ping. Is this because we are inside of this middleware function?
		return
	}
	next(w, r)
}

func main() {
	initDb()
	// template.ParseFiles() returns and error. We wrap that in template.Must() which will absorb the error from Parsefiles and halt
	// execution of the program - Is this proper error handling??
	templates := template.Must(template.ParseFiles("templates/index.html"))

	//db, _ = sql.Open("sqlite3", "dev.db") This was moved to initDb() when introducing go-gorp

	mux := gmux.NewRouter()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := Page{Books: []Book{}}
		// This is gorp Select(). It takes an object and a sql statement as its arugments
		_, err := dbmap.Select(&p.Books, "SELECT * from books")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// This code is not needed when using gorp stuff above
		// rows, _ := db.Query("SELECT pk, title, author, classification FROM books")
		// // iterate our rows object to pull the data from each book object
		// for rows.Next() {
		// 	//Create a blank book object
		// 	var b Book
		// 	// scan our the current row and place its conents in our book
		// 	// The order of the scan args must match the order in which we returned the columns in our DB query
		// 	rows.Scan(&b.PK, &b.Title, &b.Author, &b.Classification)
		// 	p.Books = append(p.Books, b)
		// }

		// ExecuteTemplate() takes the write object, the html, and a data thang and also returns an error
		// Here, we set err to the returned error object and if it is not nil then we throw the error
		err = templates.ExecuteTemplate(w, "index.html", p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		// This is what is written to the browser
		// fmt.Fprint(w, "Hello, BITCH")
	}).Methods("GET")

	mux.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		// Dummy data used prior to building search function
		// results := []SearchResult{
		// 	SearchResult{"Moby-Dick", "Herman Melville", "1851", "222222"},
		// 	SearchResult{"The Adventures of Huckleberry Finn", "Mark Twain", "1884", "444444"},
		// 	SearchResult{"The Catcher In the Rye", "JD Salinger", "1951", "333333"},
		// }
		var results []SearchResult
		var err error
		results, err = search(r.FormValue("search"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		encoder := json.NewEncoder(w)
		err = encoder.Encode(results)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}).Methods("POST")

	mux.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		var book ClassifyBookResponse
		var err error

		book, err = find(r.FormValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// This code is no longer necessary because of gorp
		// result, err := db.Exec("insert into books (pk, title, author, id, classification) values (?, ?, ?, ?, ?)",
		// 	nil, book.BookData.Title, book.BookData.Author, book.BookData.ID, book.Classification.MostPopular)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// }
		// pk, _ := result.LastInsertId()

		b := Book{
			//PK was original 'int(pk)' but gorp will populate this with a new value on insert, so we set it to -1
			PK:             -1,
			Title:          book.BookData.Title,
			Author:         book.BookData.Author,
			Classification: book.Classification.MostPopular,
		}
		// call insert to the book object we just created
		err = dbmap.Insert(&b)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(b)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}).Methods("PUT")

	mux.HandleFunc("/books/{pk}", func(w http.ResponseWriter, r *http.Request) {
		pk, _ := strconv.ParseInt(gmux.Vars(r)["pk"], 10, 64)

		// db.Exec("DELETE from books where pk = ?", gmux.Vars(r)["pk"]) was called originall before gorp
		// we call Delete on our dbmap and pass it a book pointer and leave all columns blank except our newly parsed pk int
		_, err := dbmap.Delete(&Book{pk, "", "", "", ""})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}).Methods("DELETE")

	n := negroni.Classic()
	// Use our custom middleware to check for DB connection
	n.Use(negroni.HandlerFunc(verifyDataBase))
	n.UseHandler(mux)
	// This replaces the listenAndServe
	n.Run(":8080")

}

type ClassifySearchResponse struct {
	Results []SearchResult `xml:"works>work"`
}

func find(id string) (ClassifyBookResponse, error) {
	var c ClassifyBookResponse
	body, err := classifyAPI("http://classify.oclc.org/classify2/Classify?&summary=true&owi=" + url.QueryEscape(id))
	if err != nil {
		return ClassifyBookResponse{}, err
	}
	//parse the xml response and save it to c
	err = xml.Unmarshal(body, &c)
	return c, err
}

func search(query string) ([]SearchResult, error) {
	var c ClassifySearchResponse
	body, err := classifyAPI("http://classify.oclc.org/classify2/Classify?&summary=true&title=" + url.QueryEscape(query))
	if err != nil {
		return []SearchResult{}, err
	}
	//parse the xml response and save it to c
	err = xml.Unmarshal(body, &c)
	return c.Results, err
}

func classifyAPI(url string) ([]byte, error) {
	var resp *http.Response
	var err error
	// Make the call for the data
	// url.QueryEscape(query) ensures that we have a valid http url
	resp, err = http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	// close the response once finished with everything in the function
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
