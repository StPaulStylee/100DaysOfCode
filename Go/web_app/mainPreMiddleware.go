package main

import (
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"

	_ "github.com/mattn/go-sqlite3"
)

type Page struct {
	Name     string
	DBStatus bool
}

func main() {
	// template.ParseFiles() returns and error. We wrap that in template.Must() which will absorb the error from Parsefiles and halt
	// execution of the program - Is this proper error handling??
	templates := template.Must(template.ParseFiles("templates/index.html"))

	db, _ := sql.Open("sqlite3", "dev.db")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := Page{Name: "Gopher"}
		// r.FormValue searches the query parameters for a certain value, in this case name
		// If name is unset, it will return an empty string and use the default set to "Gopher"
		name := r.FormValue("name")
		if name != "" {
			p.Name = name
		}
		p.DBStatus = db.Ping() == nil
		// ExecuteTemplate() takes the write object, the html, and a data thang and also returns an error
		// Here, we set err to the returned error object and if it is not nil then we throw the error
		err := templates.ExecuteTemplate(w, "index.html", p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		// This is what is written to the browser
		// fmt.Fprint(w, "Hello, BITCH")
	})

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
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

	http.HandleFunc("/books/add", func(w http.ResponseWriter, r *http.Request) {
		var book ClassifyBookResponse
		var err error

		book, err = find(r.FormValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		//Ping the db to check connection is alive
		err = db.Ping()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		_, err = db.Exec("insert into books (pk, title, author, id, classification) values (?, ?, ?, ?, ?)",
			nil, book.BookData.Title, book.BookData.Author, book.BookData.ID, book.Classification.MostPopular)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// we use nil here to tell the http library to use the default mux we defined above
	// We wrap the call in fmt.Println() in order to see if any errors are thrown
	fmt.Println(http.ListenAndServe(":8080", nil))
}

// These 'struct tags' (the weird `` xml thing) tells the decorder how to populate this struct from its
// XML counterpart
type SearchResult struct {
	Title  string `xml:"title,attr"`
	Author string `xml:"author,attr"`
	Year   string `xml:"hyr,attr"`
	ID     string `xml:"owi,attr"`
}

type ClassifySearchResponse struct {
	Results []SearchResult `xml:"works>work"`
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
