package main

import "fmt"

func main() {
	// var nos []int
	// var nos []int = []int{3, 1, 4, 2, 5}
	// var nos = []int{3, 1, 4, 2, 5}
	nos := []int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	// append() - to add new items to the slice
	nos = append(nos, 30)
	fmt.Println(nos)

	nos = append(nos, 10, 40, 20)
	fmt.Println(nos)

	hundreds := []int{400, 200, 100}
	nos = append(nos, hundreds...)
	fmt.Println(nos)
	fmt.Println("===============================================")

	// len() - to get the length of the slice
	fmt.Println("len(nos) =", len(nos))

	fmt.Println("Iterating the slice - [indexer]")
	for i := 0; i < len(nos); i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	fmt.Println("Iterating the slice - [range]")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}
	fmt.Println("===============================================")

	// Slices are references
	nos2 := nos // a copy of reference is assigned
	nos[0] = 9999
	fmt.Println("nos =", nos)
	fmt.Println("nos2 =", nos2)
	fmt.Println("===============================================")

	// Use copy() to create a copy
	nosCopy := make([]int, len(nos))
	copy(nosCopy, nos)
	fmt.Println("nosCopy = ", nosCopy)
	nos[0] = 8888
	fmt.Println("nos =", nos)
	fmt.Println("nosCopy = ", nosCopy)
	fmt.Println("===============================================")

	fmt.Println("[main] Before sorting, nos = ", nos)
	sortSlice(nos) // DONT need to pass a "pointer" as slices are references themselves
	fmt.Println("[main] After sorting, nos = ", nos)
	fmt.Println("===============================================")

	// slicing
	fmt.Println("nos[3:7] = ", nos[3:7])
	fmt.Println("nos[:7] = ", nos[:7])
	fmt.Println("nos[3:] = ", nos[3:])
}

func sortSlice(list []int) {
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
	fmt.Println("[sortSlice] list = ", list)
}
