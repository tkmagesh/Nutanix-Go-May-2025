/* functions as arguments */
package main

import "fmt"

func main() {
	/*
		exec("F1")
		exec("f2")
	*/

	exec(f1)
	exec(f2)
	// exec(F1) // => compilation error
	exec(func() {
		fmt.Println("anonymous fn invoked")
	})
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}

/*
func exec(fnName string) {
	// execute either f1 or f2 based on parameter
	switch fnName {
	case "f1":
		f1()
	case "f2":
		f2()
	default:
		fmt.Println("Invalid argument")
	}
}
*/

func exec(fn func()) {
	fn()
}
