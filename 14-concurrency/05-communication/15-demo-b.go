// Share memory by communicating

package main

import (
	"fmt"
)

var result int

func main() {
	ch := make(chan int)
	go add(100, 200, ch)
	result := <-ch
	fmt.Println("Result :", result)
}

func add(x, y int, ch chan int) {
	result := x + y
	ch <- result
}
