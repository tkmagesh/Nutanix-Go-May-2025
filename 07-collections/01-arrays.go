package main

import "fmt"

func main() {
	// memory is allocated and initialized with the "zero" value
	// var nos [5]int

	// initializing with data
	// var nos [5]int = [5]int{3, 1, 4, 2, 5}

	// type inference
	// var nos = [5]int{3, 1, 4, 2, 5}

	// using :=
	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	fmt.Println("Iterating the array - [indexer]")
	for i := 0; i < 5; i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	fmt.Println("Iterating the array - [range]")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	// Arrays are values
	nos2 := nos // a copy of "nos" is created
	nos[0] = 9999
	fmt.Println("nos =", nos)
	fmt.Println("nos2 =", nos2)

	x := [4]int{10, 20, 30, 40}
	y := [4]int{10, 20, 30, 40}
	fmt.Println("x == y ?", x == y) // compared by values (NOT by reference)

	// Use pointers for references
	var nosPtr *[5]int
	nosPtr = &nos
	fmt.Println((*nosPtr)[0])

	// Accessing the members of a container pointer DOES NOT require dereferencing
	fmt.Println(nosPtr[0])

	fmt.Println("[main] Before sorting, nos = ", nos)
	sortArr(&nos)
	fmt.Println("[main] After sorting, nos = ", nos)
}

func sortArr(list *[5]int) {
	for i := 0; i < (5 - 1); i++ {
		for j := i + 1; j < 5; j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
	fmt.Println("[sortArr] list = ", list)
}
