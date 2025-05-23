// Share memory by communicating

package main

import (
	"fmt"
	"sync"
)

var result int

func main() {
	wg := &sync.WaitGroup{}
	/*
		var ch chan int
		ch = make(chan int)
	*/
	ch := make(chan int)
	wg.Add(1)
	go add(100, 200, wg, ch)
	result := <-ch
	wg.Wait()
	fmt.Println("Result :", result)
}

func add(x, y int, wg *sync.WaitGroup, ch chan int) {
	result := x + y
	ch <- result
	wg.Done()
}
