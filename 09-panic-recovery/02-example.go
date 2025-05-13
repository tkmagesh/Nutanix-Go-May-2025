/*
Modify the below in such a way that the application doesn't shutdown if the user input is 0.
Instead the app should display the error and prompt the user for a new divisor again until the user is able to successfully perform divide operation.

IMPORTANT: DO NOT validate the input in the "main" function, instead react to the panic from the "divide" function appropriatively.

Except for modifying the "divide" function, feel free to make ANY change to the code

Solution: Convert the panic into an error
*/
package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divisor cannot be 0")

func main() {
	var divisor int
	for {
		fmt.Println("Enter the divisor :")
		fmt.Scanln(&divisor)
		if q, r, err := divideAdapter(100, divisor); err == nil {
			fmt.Println(q, r)
			break
		} else {
			fmt.Printf("Errror : %q, Try again!\n", err)
			continue
		}
	}
}

// adapter to conver the panic into an error
func divideAdapter(multiplier, divisor int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	quotient, remainder = divide(multiplier, divisor)
	return
}

// 3rd party library API (DO NOT change this code)
func divide(multiplier, divisor int) (quotient, remainder int) {
	// panic by custom logic
	if divisor == 0 {
		panic(ErrDivideByZero)
	}
	quotient = multiplier / divisor
	remainder = multiplier % divisor
	return
}
