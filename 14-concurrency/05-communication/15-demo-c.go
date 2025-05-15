// Share memory by communicating

package main

import (
	"fmt"
)

var result int

func main() {
	ch := make(chan int)
	go func() {
		result := add(100, 200)
		ch <- result
	}()
	result := <-ch
	fmt.Println("Result :", result)
}

func add(x, y int) int {
	result := x + y
	return result
}
