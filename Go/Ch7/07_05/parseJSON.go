package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
)

type Tour struct {
	Name, Price string
}

func main() {

	url := "http://services.explorecalifornia.org/json/tours.php"
	content := contentFromServer(url)

	// fmt.Println(content)

	tours := toursFromJson(content)
	// fmt.Print(tours)

	// This code below is to help format the JSON into a better format
	// This for loop is ignoring the index, hence the '_'
	for _, tour := range tours {
		price, _, _ := big.ParseFloat(tour.Price, 10, 2, big.ToZero) // value to parse, base 10, precision, rounding mode
		fmt.Printf("%v ($%.2f)\n", tour.Name, price)                 // format the price into dollars with 2 decimals from a float
	}

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func contentFromServer(url string) string {

	resp, err := http.Get(url)
	checkError(err)

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	return string(bytes)
}

func toursFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20) // Make a slice of tour objects, initial size of slice, maximum size of slice
	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token() // This removies the array bracket at the beginning of the json string, we only want objects
	checkError(err)

	var tour Tour
	for decoder.More() { // Is there more data to read, if so, go to it - kinda like .hasNext()
		err := decoder.Decode(&tour) // This is where the json decoding actually happens. The pointer reuses the same memory in every loop and that is why the tour object can be outside of the loop
		checkError(err)
		tours = append(tours, tour) // Append into our tours slice the current tour
	}
	return tours
}
