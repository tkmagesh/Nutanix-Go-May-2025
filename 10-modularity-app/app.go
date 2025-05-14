package main

import (
	"fmt"

	"github.com/fatih/color"
	calc "github.com/tkmagesh/Nutanix-Go-May-2025/10-modularity-app/calculator"
	"github.com/tkmagesh/Nutanix-Go-May-2025/10-modularity-app/calculator/utils"
)

func run() {
	// fmt.Println("app executed")
	color.Red("app executed")
	fmt.Println(calc.Add(100, 200))
	fmt.Println(calc.Subtract(100, 200))
	fmt.Println(calc.GetOpCount())
	fmt.Println("Is 21 even ?:", utils.IsEven(21))

}
