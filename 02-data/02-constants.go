package main

import "fmt"

func main() {

	// const pi float64 = 3.14

	// type inference
	const pi = 3.14

	// cannot re-assign a const
	// pi = 2

	// constants can remain unused
	fmt.Println(pi)

	// group const declarations
	/*
		const (
			x int     = 100
			y float64 = 20.99
		)
	*/

	// type inference
	const (
		x = 100
		y = 20.99
	)

}
