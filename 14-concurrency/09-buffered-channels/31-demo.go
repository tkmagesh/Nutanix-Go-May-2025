package main

import "fmt"

func main() {
	// ch := make(chan int) //=> deadlock
	ch := make(chan int, 1) //=> allocating memory in the channel to hold some data
	ch <- 100
	data := <-ch
	fmt.Println(data)

}
