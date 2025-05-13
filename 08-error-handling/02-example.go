/*
"Typed" error
*/
package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divisor cannot be 0")

func main() {
	// divisor := 7
	divisor := 0
	q, r, err := divide(100, divisor)
	if err == nil {
		fmt.Printf("Dividing 100 by %d, quotient = %d, remainder = %d\n", divisor, q, r)
	} else if err == ErrDivideByZero {
		fmt.Println("Do not attempt to divide by zero")
	} else {
		fmt.Println("Unknown Error occurred :", err)
	}
}

// Named results
func divide(multiplier, divisor int) (quotient, remainder int, err error) {
	if divisor == 0 {
		err = ErrDivideByZero
		return
	}
	quotient, remainder = multiplier/divisor, multiplier%divisor
	return
}
