package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	sync.Mutex
	count int
}

func (sf *SafeCounter) Add(delta int) {
	sf.Lock()
	{
		sf.count += delta
	}
	sf.Unlock()
}

var sf SafeCounter

func main() {
	wg := &sync.WaitGroup{}
	for range 300 {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println("Count :", sf.count)
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	sf.Add(1)
}
