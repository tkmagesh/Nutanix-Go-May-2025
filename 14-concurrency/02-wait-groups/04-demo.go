package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	for id := range 20 {
		wg.Add(1)     //increment the counter by 1
		go fn(wg, id) //scheduling the execution of f1 through the scheduler
	}
	wg.Wait() // block the execution of this function until the counter becomes 0 (default = 0)
	fmt.Println("Done!!")
}

func fn(wg *sync.WaitGroup, id int) {
	defer wg.Done() // decrement the counter by 1
	fmt.Printf("fn[%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	fmt.Printf("fn[%d] completed\n", id)
}
