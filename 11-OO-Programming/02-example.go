package main

import "fmt"

type Product struct {
	id   int
	name string
	cost float64
}

func main() {
	// product -> id, name, cost
	/*
		var p Product
		p.id = 100
		p.name = "Pen"
		p.cost = 10
	*/
	var p Product = Product{
		id:   100,
		name: "pen",
		cost: 10,
	}
	// fmt.Println(p)
	// fmt.Printf("%#v\n", p)
	// fmt.Printf("%+v\n", p)
	fmt.Println(Format(p))

	fmt.Println("Before applying discount :", Format(p))
	ApplyDiscount(/* ??? */) // 10%
	fmt.Println("After applying discount :", Format(p))
}

func Format(p Product) string {
	return fmt.Sprintf("id = %d, name = %q, cost = %0.2f", p.id, p.name, p.cost)
}

// write a function that applies the given discount (percentage) to the given product
func ApplyDiscount(/* product */, /* discountPercentage */) /* do not return */ {
	/* logic to apply the discount */
}
