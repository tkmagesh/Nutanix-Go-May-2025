/* functions as return values */

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fn := getFn()
	fn()
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}

func getFn() func() {
	if no := rand.Intn(20); no%2 == 0 {
		return f1
	}
	return f2
}
