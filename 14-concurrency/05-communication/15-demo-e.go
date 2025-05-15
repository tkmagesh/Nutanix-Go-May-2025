// Share memory by communicating

package main

import (
	"fmt"
	"time"
)

var result int

// consumer
func main() {

	ch := add(100, 200)
	result := <-ch
	fmt.Println("Result :", result)
}

// producer
func add(x, y int) <-chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(500 * time.Millisecond)
		result := x + y
		ch <- result
	}()
	return ch
}
