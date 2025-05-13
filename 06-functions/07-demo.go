package main

import "fmt"

func main() {
	var n int = 100
	fmt.Println("[main] Address  :", &n)
	fmt.Printf("[main] Before incrementing, n : %d\n", n)
	increment(&n)
	fmt.Printf("[main] After incrementing, n : %d\n", n)

}

func increment(no *int) {
	fmt.Println("[increment] Address  :", no)
	*no += 1
	fmt.Printf("[increment] no : %d\n", *no)
}
