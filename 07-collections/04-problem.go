/*
Refactor the below as follows:
genPrimes(start, end) returns the list of prime numbers between the "start" and the "end"
Print each prime no in the main() function
*/
package main

import "fmt"

func main() {
LOOP:
	for no := 2; no <= 100; no++ {
		for i := 2; i <= (no / 2); i++ {
			if no%i == 0 {
				continue LOOP
			}
		}
		fmt.Println("Prime No :", no)
	}
}
