package main

import "fmt"

func main() {
	var no int
	no = 100

	no2 := no // get a copy of the value
	no2 += 10 // will not affect the "no"

	// noPtr := &no // get the address of no
	var noPtr *int
	noPtr = &no
	fmt.Println(noPtr)
	fmt.Printf("Type of noPtr : %T\n", noPtr)

	fmt.Println("Data at noPtr :", *noPtr)

	no2Ptr := &no
	*no2Ptr += 10
	fmt.Printf("no : %d, no2Ptr : %d\n", no, *no2Ptr)
}
