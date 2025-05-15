/*
Refactor the following so that the "printing of the prime numbers" happen in the main function
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go generatePrimes(2, 25, wg)
	wg.Add(1)
	go generatePrimes(26, 50, wg)
	wg.Add(1)
	go generatePrimes(51, 75, wg)
	wg.Add(1)
	go generatePrimes(76, 100, wg)
	wg.Wait()
	fmt.Println("Done!")
}

func generatePrimes(start, end int, wg *sync.WaitGroup) {
	defer wg.Done()
LOOP:
	for no := start; no <= end; no++ {
		for i := 2; i <= (no / 2); i++ {
			if no%i == 0 {
				continue LOOP
			}
		}
		fmt.Println("Prime No :", no)
	}

}
