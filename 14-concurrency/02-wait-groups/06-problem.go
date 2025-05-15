/*
Execute the logic for finding & printing prime numbers concurrently
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
