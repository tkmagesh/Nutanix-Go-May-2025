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

// maintainable
func genNos(wg *sync.WaitGroup) {
	defer wg.Done()
	timeoutCh := time.After(10 * time.Second)
LOOP:
	for i := 1; ; i++ {
		select {
		case <-timeoutCh:
			break LOOP
		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Println(i)
		}
	}
	fmt.Println("[genNos] stopped")
}

/* func timeout(d time.Duration) <-chan time.Time {
	timeoutCh := make(chan time.Time)
	go func() {
		time.Sleep(d)
		timeoutCh <- time.Now()
	}()
	return timeoutCh
} */
