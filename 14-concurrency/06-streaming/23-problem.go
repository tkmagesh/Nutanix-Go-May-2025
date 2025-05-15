/*
Update the following using "Share memory by communicating" (channels)
*/

package main

import (
	"fmt"
	"sync"
)

var primes []int
var mutex sync.Mutex

func main() {
	ch := make(chan int)
	go func() {
		dataCh := make(chan int)
		go func() {
			for no := 2; no <= 100; no++ {
				dataCh <- no
			}
			close(dataCh)
		}()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go generatePrimes(dataCh, wg, ch)
		wg.Add(1)
		go generatePrimes(dataCh, wg, ch)
		wg.Add(1)
		go generatePrimes(dataCh, wg, ch)
		wg.Add(1)
		go generatePrimes(dataCh, wg, ch)
		wg.Wait()
		close(ch)
	}()
	for primeNo := range ch {
		fmt.Printf("Prime No : %d\n", primeNo)
	}
	fmt.Println("Done!")
}

func generatePrimes(dataCh <-chan int, wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
LOOP:
	for no := range dataCh {
		for i := 2; i <= (no / 2); i++ {
			if no%i == 0 {
				continue LOOP
			}
		}
		ch <- no
	}
}
