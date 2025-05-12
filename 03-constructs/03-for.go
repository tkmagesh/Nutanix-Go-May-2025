package main

import "fmt"

func main() {
	// Ver 1.0
	fmt.Println("Ver 1.0")
	for i := 0; i < 10; i++ {
		fmt.Printf("i = %d\n", i)
	}

	fmt.Println("Ver 2.0 - while")
	i := 0
	for i < 10 {
		fmt.Printf("i = %d\n", i)
		i++
	}

	fmt.Println("Ver 3.0")
	for i := 0; i < 10; {
		fmt.Printf("i = %d\n", i)
		i++
	}

	fmt.Println("Ver 4.0")
	for i := 0; ; i++ {
		fmt.Printf("i = %d\n", i)
		if i >= 10 {
			break
		}
	}

	fmt.Println("Ver 5.0")
	no := 0
	for {
		fmt.Printf("no = %d\n", no)
		no++
		if no >= 10 {
			break
		}
	}

	fmt.Println("Ver 6.0")
	for i := range 10 {
		fmt.Println(i)
	}

	fmt.Println("Ver 7.0")
	for range 10 {
		fmt.Println("do something")
	}

	// Using labels
	fmt.Println("Using labels")
X_LOOP:
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			fmt.Printf("x = %d, y = %d\n", x, y)
			if x == y {
				fmt.Println("====================")
				continue X_LOOP // controlling the outer loop from an inner loop using label
				// goto X_LOOP
			}
		}
	}
}
