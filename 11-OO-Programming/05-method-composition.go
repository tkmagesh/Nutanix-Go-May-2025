package main

import (
	"fmt"
	"strconv"
)

type Product struct {
	id   int
	name string
	cost float64
}

// Method => function with a receiver
func (p Product) WhoAmI() string {
	return "I am a product : " + strconv.Itoa(p.id)
}
func (p Product) Format() string {
	return fmt.Sprintf("id = %d, name = %q, cost = %0.2f", p.id, p.name, p.cost)
}

func (p *Product) ApplyDiscount(discountPercentage float64) /* do not return */ {
	p.cost = p.cost * ((100 - discountPercentage) / 100)
}

// Make the following work for a PerishableProduct
func main() {

	var p Product = Product{
		id:   100,
		name: "pen",
		cost: 10,
	}
	fmt.Println(p.WhoAmI())

	fmt.Println("Before applying discount :", p.Format())

	p.ApplyDiscount(10)

	fmt.Println("After applying discount :", p.Format())

}

// write a function that applies the given discount (percentage) to the given product
