package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go genNos(wg)
	wg.Wait()
}

// unmaintainable
func genNos(wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	for i := 1; ; i++ {
		elapsed := time.Since(start)
		if elapsed >= time.Duration(10*time.Second) {
			break
		}
		time.Sleep(500 * time.Millisecond)
		fmt.Println(i)
	}
	fmt.Println("[genNos] stopped")
}
