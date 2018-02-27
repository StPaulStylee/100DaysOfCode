package main

import (
	"fmt"
	"net/http"
)

type Hello struct{}

// This function must be named ServeHTTP() and these argument methods must be in this order and *http.Request must be a reference
func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint prints to a writer object, which has been passed in with the 'w'
	fmt.Fprint(w, "<h1> Hello from the Go web server!<h1>")
}

func main() {
	var h Hello
	// ListenAndServce takes a url and port number and the second is an instance of the Hello object...
	// ... This is where the method will search and call the ServeHTTP method
	err := http.ListenAndServe("localhost:5050", h)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
