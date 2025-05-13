package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n1, n2 int
	fmt.Println("Enter the numbers :")
	fmt.Scanln(&n1, &n2)

	// decide the operation to be performed
	var op func(int, int)
	op = getOp()

	// perform the operation
	op(n1, n2)

}

func getOp() func(int, int) {
	var op func(int, int)
	r := rand.Intn(20)
	switch {
	case r%2 == 0:
		op = func(i1, i2 int) {
			fmt.Println("Add Result :", i1+i2)
		}
	case r%3 == 0:
		op = func(i1, i2 int) {
			fmt.Println("Subtract Result :", i1-i2)
		}
	case r%5 == 0:
		op = func(i1, i2 int) {
			fmt.Println("Multiply Result :", i1*i2)
		}
	default:
		op = func(i1, i2 int) {
			fmt.Println("Divide Result :", i1/i2)
		}
	}
	return op
}
