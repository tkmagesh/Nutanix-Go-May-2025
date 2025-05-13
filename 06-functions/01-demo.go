package main

import "fmt"

func main() {
	sayHi()
	sayHello("Magesh")
	greet("Magesh", "Kuppan")
	fmt.Println(add(100, 200))
	// fmt.Println(divide(100, 7))

	q, r := divide(100, 7)
	fmt.Printf("dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)

	// Use only quotient
	// Use "_" in place of a variable to RECEIVE a value that is not used

	/*
		q, _ := divide(100, 7)
		fmt.Printf("dividing 100 by 7, quotient = %d \n", q)
	*/

}

func sayHi() {
	fmt.Println("Hi there!")
}

func sayHello(userName string) {
	fmt.Printf("Hi %s, Have a nice day!\n", userName)
}

/*
func greet(firstName string, lastName string) {
	fmt.Printf("Hi %s %s, Have a good day!\n", firstName, lastName)
}
*/

func greet(firstName, lastName string) {
	fmt.Printf("Hi %s %s, Have a good day!\n", firstName, lastName)
}

func add(x, y int) int {
	return x + y
}

// more than one return result
func divide(x, y int) (int, int) {
	var quotient, remainder int
	quotient = x / y
	remainder = x % y
	return quotient, remainder
}

// more than one return result (using NAMED results) (ALWAYS use "named" results when there are more than one return results)
/* func divide(x, y int) (quotient, remainder int) {
	quotient = x / y
	remainder = x % y
	return quotient, remainder
} */
