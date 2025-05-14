package main

import (
	"fmt"

	"github.com/tkmagesh/Nutanix-Go-May-2025/10-modularity-app/calculator"
	"github.com/tkmagesh/Nutanix-Go-May-2025/10-modularity-app/calculator/utils"
)

func run() {
	fmt.Println("app executed")
	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))
	fmt.Println(calculator.GetOpCount())
	fmt.Println("Is 21 even ?:", utils.IsEven(21))

}
