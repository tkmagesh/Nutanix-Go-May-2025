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
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go generatePrimes(2, 25, wg, ch)
		wg.Add(1)
		go generatePrimes(26, 50, wg, ch)
		wg.Add(1)
		go generatePrimes(51, 75, wg, ch)
		wg.Add(1)
		go generatePrimes(76, 100, wg, ch)
		wg.Wait()
		close(ch)
	}()
	for primeNo := range ch {
		fmt.Printf("Prime No : %d\n", primeNo)
	}
	fmt.Println("Done!")
}

func generatePrimes(start, end int, wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
LOOP:
	for no := start; no <= end; no++ {
		for i := 2; i <= (no / 2); i++ {
			if no%i == 0 {
				continue LOOP
			}
		}
		ch <- no
	}
}
