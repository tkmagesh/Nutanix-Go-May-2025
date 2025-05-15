// Share memory by communicating

package main

import (
	"fmt"
)

var result int

// consumer
func main() {

	ch := add(100, 200)
	result := <-ch
	fmt.Println("Result :", result)
}

// producer
func add(x, y int) chan int {
	ch := make(chan int)
	result := x + y
	ch <- result
	return ch
}
