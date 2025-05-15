/*
Execute the logic for finding & printing prime numbers concurrently
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go generatePrimes(2, 100, wg)
	wg.Wait()
	fmt.Println("Done!")
}

func generatePrimes(start, end int, wg *sync.WaitGroup) {
	defer wg.Done()
	wg2 := &sync.WaitGroup{}
	for no := start; no <= end; no++ {
		wg2.Add(1)
		go check_and_print_prime(no, wg2)
	}
	wg2.Wait()
}

func check_and_print_prime(no int, wg2 *sync.WaitGroup) {
	defer wg2.Done()
	if isPrime(no) {
		fmt.Printf("Prime No : %d\n", no)
	}
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
