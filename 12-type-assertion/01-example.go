package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var x interface{}
	x = 100
	x = true
	x = "Aliquip laborum tempor Lorem laboris magna anim culpa reprehenderit nulla quis anim ipsum incididunt amet."
	x = 99.98
	x = struct{}{}
	fmt.Println(x)

	// x = 200
	x = getSomeValue()
	// y := x * 2

	// unsafe
	// y := x.(int) * 2

	// safe (type assertion - using if-else)
	if val, ok := x.(int); ok {
		y := val * 2
		fmt.Println(y)
	} else {
		fmt.Println("x is not an int")
	}

	// safe (type assertion - using "type" switch)
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x * 2 =", val*2)
	case string:
		fmt.Println("x is a string, len(x) =", len(val))
	case bool:
		fmt.Println("x is a bool, !x =", !val)
	case float64:
		fmt.Println("x is a float64, 0.9 * x =", 0.9*val)
	default:
		fmt.Println("x is of unknown type")
	}
}

func getSomeValue() interface{} {
	if rand.Intn(20)%2 == 0 {
		return "Non anim officia reprehenderit aliqua aute cillum aliqua proident id aliquip deserunt laborum."
	}
	return 200
}
