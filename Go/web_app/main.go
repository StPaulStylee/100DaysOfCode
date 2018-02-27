package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// This is what is written to the browser
		fmt.Fprint(w, "Hello, BITCH")
	})
	// we use nil here to tell the http library to use the default mux we defined above
	// We wrap the call in fmt.Println() in order to see if any errors are thrown
	fmt.Println(http.ListenAndServe(":8080", nil))
}
