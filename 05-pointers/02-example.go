package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	/*
		var userName string
		fmt.Println("Enter your name :")
		fmt.Scanln(&userName)
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	*/

	/*
		var n1, n2 int
		fmt.Println("Enter 2 numbers :")
		fmt.Scanln(&n1, &n2)
		fmt.Printf("n1 = %d, n2 = %d\n", n1, n2)
	*/

	// Accepting data with spaces

	var scanner *bufio.Scanner
	scanner = bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a sentence:")
	scanner.Scan()
	txt := scanner.Text()
	fmt.Println(strings.ToUpper(txt))

}
