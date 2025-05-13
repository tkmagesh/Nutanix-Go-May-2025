package main

import (
	"fmt"
)

func main() {
	// var productRanks map[string]int
	// var productRanks map[string]int = map[string]int{"Pen": 3, "Pencil": 1, "Marker": 5}
	// var productRanks = map[string]int{"Pen": 3, "Pencil": 1, "Marker": 5}
	// productRanks := map[string]int{"Pen": 3, "Pencil": 1, "Marker": 5}

	/*
		// Declaration
		var productRanks map[string]int

		// Initialization
		// productRanks = make(map[string]int)
		productRanks = map[string]int{}
	*/

	// Declaration & Initialization
	productRanks := make(map[string]int)
	productRanks["Pen"] = 3
	productRanks["Pencil"] = 1
	productRanks["Marker"] = 5
	fmt.Println(productRanks)
	fmt.Println("=========================================================")

	fmt.Printf("len(productRanks) : %d\n", len(productRanks))
	fmt.Println("=========================================================")

	fmt.Println("Iterating a map")
	for key, val := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, val)
	}
	fmt.Println("=========================================================")

	// check for the existence of a key
	keyToCheck := "Scribble-Pad"
	// keyToCheck := "Marker"

	if r, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("Rank of %q is %d\n", keyToCheck, r)
	} else {
		fmt.Printf("key - %q does not exist\n", keyToCheck)
	}

	fmt.Println("=========================================================")
	// Removing an item
	// keyToRemove := "Pen"
	keyToRemove := "Scribble-Pad"
	delete(productRanks, keyToRemove)
	fmt.Printf("After removing key - %q, productRanks :\n", keyToRemove)
	fmt.Println(productRanks)

}
