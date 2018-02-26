package main

import (
	"fmt"
	"sort"
)

func main() {
	// Always wrap your maps with the make function otherwise you won't be able to assign any values to it
	// because the memory is essentially "locked". This is initializing a map with keys that are strings and values that are also strings
	states := make(map[string]string)
	fmt.Println(states)

	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"
	fmt.Println(states)

	California := states["CA"]
	fmt.Println(California)

	delete(states, "OR")
	fmt.Println(states)

	states["NY"] = "New York"

	// This is how you can iterate over the key/value pairs in a map
	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	// To set the order of the map to alphabetical...
	// First initialize a slice that is the size of the states map
	keys := make([]string, len(states))
	i := 0
	// When iterating with the range keyword and you only assign one variable, you get the key and not its associated value
	// Here it loops through the 'states' map and assigns the current index of the slice with the key
	for k := range states {
		keys[i] = k
		i++
	}
	// Then we sort the keys into alphabetical order
	sort.Strings(keys)
	fmt.Println("\nSorted")
	// and loop through the keys slice and print the states map key values using the keys from the slice... Crazy.
	for i := range keys {
		fmt.Println(states[keys[i]])
	}

}
