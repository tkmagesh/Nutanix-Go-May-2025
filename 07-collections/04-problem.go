/*
Refactor the below as follows:
genPrimes(start, end) returns the list of prime numbers between the "start" and the "end"
Print each prime no in the main() function
*/
package main

import "fmt"

func main() {

	primes := genPrimes(2, 100)
	for _, primeNo := range primes {
		fmt.Printf("Prime No : %d\n", primeNo)
	}
}

func genPrimes(start, end int) []int {
	var primes []int
	for no := start; no <= end; no++ {
		if isPrime(no) {
			primes = append(primes, no)
		}
	}
	return primes
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
