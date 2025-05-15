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

/*
func (p Product) Format() string {
	return fmt.Sprintf("id = %d, name = %q, cost = %0.2f", p.id, p.name, p.cost)
}
*/

// fmt.Stringer interface implementation
func (p Product) String() string {
	return fmt.Sprintf("id = %d, name = %q, cost = %0.2f", p.id, p.name, p.cost)
}

func (p *Product) ApplyDiscount(discountPercentage float64) /* do not return */ {
	p.cost = p.cost * ((100 - discountPercentage) / 100)
}

// PerishableProduct
type PerishableProduct struct {
	Product
	expiry string
}

func (pp PerishableProduct) WhoAmI() string {
	return "I am a perishable product : " + strconv.Itoa(pp.id)
}

/*
func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%s, expiry = %q", pp.Product.Format(), pp.expiry)
}
*/

// fmt.Stringer interface implementation
func (pp PerishableProduct) String() string {
	return fmt.Sprintf("%s, expiry = %q", pp.Product, pp.expiry)
}

// factory function for PerishableProduct
func NewPerishableProduct(id int, name string, cost float64, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			id:   id,
			name: name,
			cost: cost,
		},
		expiry: expiry,
	}
}

// Make the following work for a PerishableProduct
func main() {

	/*
		var milk = PerishableProduct{
			Product: Product{
				id:   101,
				name: "Milk",
				cost: 50,
			},
			expiry: "2 Days",
		}
	*/
	var milk = NewPerishableProduct(102, "Milk", 50, "2 Days")

	fmt.Println(milk.WhoAmI())

	fmt.Println("Before applying discount :", milk)

	milk.ApplyDiscount(10)

	fmt.Println("After applying discount :", milk)

}

// write a function that applies the given discount (percentage) to the given product
