package main

import (
	"fmt"
	"time"
)

func main() {
	ch := genNos()
	for range 50 {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(<-ch)
	}
	fmt.Println("Done!")
}

func genNos() <-chan int {
	ch := make(chan int)
	go func() {
		for i := range 5 {
			ch <- (i + 1) * 10
		}
		close(ch)
	}()
	return ch
}
