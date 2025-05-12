package main

import "fmt"

func main() {

	/*
		const Red int = 0
		const Green int = 1
		const Blue int = 2
		fmt.Printf("Red = %d, Green = %d and Blue = %d\n", Red, Green, Blue)
	*/

	// auto generating constant values

	/*
		const (
			Red   int = iota
			Green int = iota
			Blue  int = iota
		)
		fmt.Printf("Red = %d, Green = %d and Blue = %d\n", Red, Green, Blue)
	*/

	/*
		const (
			Red int = iota
			Green
			Blue
		)
		fmt.Printf("Red = %d, Green = %d and Blue = %d\n", Red, Green, Blue)
	*/

	// type inference
	/*
		const (
			Red = iota
			Green
			Blue
		)
		fmt.Printf("Red = %d, Green = %d and Blue = %d\n", Red, Green, Blue)
	*/

	/*
		const (
			Red   = iota + 2 // 0 + 2
			Green            // 1 + 2
			Blue             // 2 + 2
		)
		fmt.Printf("Red = %d, Green = %d and Blue = %d\n", Red, Green, Blue)
	*/
	/*
		const (
			Red   = (iota * 2) + 1 // (0 * 2) + 1
			Green                  // (1 * 2) + 1
			Blue                   // (2 * 2) + 1
		)
		fmt.Printf("Red = %d, Green = %d and Blue = %d\n", Red, Green, Blue)
	*/

	// skip a value
	/*
		const (
			Red   = (iota * 2) + 1 // (0 * 2) + 1
			Green                  // (1 * 2) + 1
			_
			Blue // (3 * 2) + 1
		)
		fmt.Printf("Red = %d, Green = %d and Blue = %d\n", Red, Green, Blue)
	*/

	// Mimicking Unix file permissions

	/*
		const (
			X = iota << 1
			W
			R
		)
		fmt.Println(X, W, R)
		fmt.Printf("%b %b %b\n", X, W, R) // 1, 10, 100
	*/

	const (
		X   = 1 << iota // 1 << 0 == 001
		W               // 1 << 1 == 010
		R               // 1 << 2 == 100
		XW  = X | W
		WR  = W | R
		XWR = X | W | R
	)
	fmt.Printf("%b %b %b %b %b %b\n", X, W, XW, R, WR, XWR)

}
