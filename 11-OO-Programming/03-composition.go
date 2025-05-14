package main

import "fmt"

type Product struct {
	id   int
	name string
	cost float64
}

type Dummy struct {
	id int
}

type PerishableProduct struct {
	Product
	// Dummy
	expiry string
}

func main() {
	milk := PerishableProduct{
		expiry: "2 Days",
		Product: Product{
			id:   101,
			name: "Milk",
			cost: 50,
		},
	}
	fmt.Println(milk)

	// Accessing the attributes
	fmt.Println(milk.expiry)
	/*
		fmt.Println(milk.Product.id)
		fmt.Println(milk.Product.name)
		fmt.Println(milk.Product.cost)
	*/
	fmt.Println(milk.id)
	fmt.Println(milk.name)
	fmt.Println(milk.cost)

}
