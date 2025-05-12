package main

import "fmt"

func main() {
	var i int8 = 10 // 1 byte
	var f float64   // 8 bytes

	// f = i // compilation error

	// use the typename like a function for type conversion
	f = float64(i)
	fmt.Println(f)
}
