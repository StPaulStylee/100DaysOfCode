package main

import (
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/codegangsta/negroni"
	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	PK             int
	Title          string
	Author         string
	Classification string
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

func main() {
	// template.ParseFiles() returns and error. We wrap that in template.Must() which will absorb the error from Parsefiles and halt
	// execution of the program - Is this proper error handling??
	templates := template.Must(template.ParseFiles("templates/index.html"))

	db, _ = sql.Open("sqlite3", "dev.db")

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := Page{Books: []Book{}}
		rows, _ := db.Query("SELECT pk, title, author, classification FROM books")
		// iterate our rows object to pull the data from each book object
		for rows.Next() {
			//Create a blank book object
			var b Book
			// scan our the current row and place its conents in our book
			// The order of the scan args must match the order in which we returned the columns in our DB query
			rows.Scan(&b.PK, &b.Title, &b.Author, &b.Classification)
			p.Books = append(p.Books, b)
		}

		// ExecuteTemplate() takes the write object, the html, and a data thang and also returns an error
		// Here, we set err to the returned error object and if it is not nil then we throw the error
		err := templates.ExecuteTemplate(w, "index.html", p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		// This is what is written to the browser
		// fmt.Fprint(w, "Hello, BITCH")
	})

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
	})

	mux.HandleFunc("/books/add", func(w http.ResponseWriter, r *http.Request) {
		var book ClassifyBookResponse
		var err error

		book, err = find(r.FormValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		result, err := db.Exec("insert into books (pk, title, author, id, classification) values (?, ?, ?, ?, ?)",
			nil, book.BookData.Title, book.BookData.Author, book.BookData.ID, book.Classification.MostPopular)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		pk, _ := result.LastInsertId()
		b := Book{
			PK:             int(pk),
			Title:          book.BookData.Title,
			Author:         book.BookData.Author,
			Classification: book.Classification.MostPopular,
		}
		err = json.NewEncoder(w).Encode(b)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("/books/delete", func(w http.ResponseWriter, r *http.Request) {
		_, err := db.Exec("DELETE from books where pk = ?", r.FormValue("pk"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

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
