package main

import "fmt"

// unused variable (package scope - allowed)
var x int

func main() {

	/*
		var userName string
		userName = "Magesh"
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	*/
	/*
		var userName string = "Magesh"
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	*/

	// type inference (by the compiler)

	/*
		var userName = "Magesh"
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	*/

	// Using ":=" (shortcut for "declaration + initialization")

	userName := "Magesh"
	fmt.Printf("Hi %s, Have a nice day!\n", userName)

	// unused variable (function scope - not allowed)
	/*
		var x int
		x = 100
		fmt.Println(x)
	*/

	// Multiple variables

	/*
		var x int
		var y int
		var str string
		var result int
		x = 100
		y = 200
		str = "sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	// combining variable declarations of same type
	/*
		var x, y, result int
		var str string
		x = 100
		y = 200
		str = "sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	// combining variable declarations of different types
	/*
		var (
			x, y, result int
			str          string
		)
		x = 100
		y = 200
		str = "sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	// declaration & initialization
	/*
		var x int = 100
		var y int = 200
		var str string = "sum of %d and %d is %d\n"
		var result int = x + y
		fmt.Printf(str, x, y, result)
	*/

	// multi assignment
	/*
		var x,y int = 100, 200
		var str string = "sum of %d and %d is %d\n"
		var result int = x + y
		fmt.Printf(str, x, y, result)
	*/

	// combining variable declarations & initialization of different types
	/*
		var (
			x, y   int    = 100, 200
			str    string = "sum of %d and %d is %d\n"
			result int    = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	// type inference
	/*
		var (
			x, y   = 100, 200
			str    = "sum of %d and %d is %d\n"
			result = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y, str = 100, 200, "sum of %d and %d is %d\n"
			result    = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	// using ":="
	/*
		x, y, str := 100, 200, "sum of %d and %d is %d\n"
		result := x + y
		fmt.Printf(str, x, y, result)
	*/

}
