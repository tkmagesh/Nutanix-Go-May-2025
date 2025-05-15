/*
Update the following using "Share memory by communicating" (channels)
*/

package main

import (
	"flag"
	"fmt"
	"sync"
)

var primes []int
var mutex sync.Mutex

func main() {
	var workerCount int
	ch := make(chan int)
	flag.IntVar(&workerCount, "count", 1, "Number of workers to use")
	flag.Parse()
	go func() {
		dataCh := make(chan int)
		go func() {
			for no := 2; no <= 100; no++ {
				dataCh <- no
			}
			close(dataCh)
		}()
		wg := &sync.WaitGroup{}
		for workerId := range workerCount {
			fmt.Printf("Starting worker  : %d\n", workerId)
			wg.Add(1)
			go generatePrimes(workerId, dataCh, wg, ch)
		}
		wg.Wait()
		close(ch)
	}()
	for primeNo := range ch {
		fmt.Printf("Prime No : %d\n", primeNo)
	}
	fmt.Println("Done!")
}

func generatePrimes(workerId int, dataCh <-chan int, wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
LOOP:
	for no := range dataCh {
		fmt.Printf("Worker #%d processing %d\n", workerId, no)
		for i := 2; i <= (no / 2); i++ {
			if no%i == 0 {
				continue LOOP
			}
		}
		ch <- no
	}
	fmt.Printf("Worker #%d shutting down\n", workerId)
}
