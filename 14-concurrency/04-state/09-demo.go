package main

import (
	"fmt"
	"sync"
)

var count int

func main() {
	wg := &sync.WaitGroup{}
	for range 300 {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println("Count :", count)
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	count += 1
}
