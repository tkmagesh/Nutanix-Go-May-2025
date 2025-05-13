package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divisor cannot be 0")

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("[main - deferred] - application shutdown due to error:", err)
			return
		}
		fmt.Println("[main - deferred] - Thank You!")
	}()

	// divisor := 7
	divisor := 0
	q, r := divide(100, divisor)
	fmt.Printf("Dividing 100 by %d, quotient = %d, remainder = %d\n", divisor, q, r)

}

func divide(multiplier, divisor int) (quotient, remainder int) {
	defer func() {
		fmt.Println(" [divide] - deferred")
	}()
	fmt.Println("[divide] - calculating quotient")
	// quotient = multiplier / divisor //=> resulted in a "builtin" panic

	// panic by custom logic
	if divisor == 0 {
		panic(ErrDivideByZero)
	}

	quotient = multiplier / divisor
	fmt.Println("[divide] - calculating remainder")
	remainder = multiplier % divisor
	return
}
