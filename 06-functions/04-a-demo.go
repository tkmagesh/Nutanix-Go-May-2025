/*
	 Higher Order Functions
		- Assign functions as values to variables
*/
package main

import "fmt"

func main() {
	var sayHi func()
	sayHi = func() {
		fmt.Println("Hi there!")
	}
	sayHi()

	// var sayHello func(string)
	var sayHello func(userName string)
	sayHello = func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}
	sayHello("Magesh")

	var greet func(string, string)
	greet = func(firstName, lastName string) {
		fmt.Printf("Hi %s %s, Have a good day!\n", firstName, lastName)
	}
	greet("Magesh", "Kuppan")

	greet = func(firstName, lastName string) {
		fmt.Printf("Hi %s %s, Have a wonderful day!\n", firstName, lastName)
	}
	greet("Ramesh", "Jayaraman")

	var add func(int, int) int
	add = func(x, y int) int {
		return x + y
	}
	addResult := add(100, 200)
	fmt.Println("Add Result :", addResult)

	var divide func(int, int) (int, int)
	divide = func(x, y int) (quotient, remainder int) {
		quotient, remainder = x/y, x%y
		return
	}
	q, r := divide(100, 7)
	fmt.Printf("dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)

}
