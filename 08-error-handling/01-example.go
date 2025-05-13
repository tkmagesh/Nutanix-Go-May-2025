package main

import (
	"errors"
	"fmt"
)

func main() {
	// divisor := 0
	divisor := 7
	if q, r, err := divide(100, divisor); err == nil {
		fmt.Printf("dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
	} else {
		fmt.Println("Error occurred :", err)
	}
}

func divide(multiplier, divisor int) (int, int, error) {
	if divisor == 0 {
		var err error
		err = errors.New("divisor cannot be 0")
		return 0, 0, err
	}
	quotient, remainder := multiplier/divisor, multiplier%divisor
	return quotient, remainder, nil
}

func divide(multiplier, divisor int) (quotient, remainder int, err error) {
	if divisor == 0 {

		err = errors.New("divisor cannot be 0")
		return
	}
	quotient, remainder = multiplier/divisor, multiplier%divisor
	return
}
