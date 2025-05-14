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
	ApplyDiscount(&p, 10) // 10%
	fmt.Println("After applying discount :", Format(p))

	p1 := Product{101, "Pencil", 5}
	p2 := Product{101, "Pencil", 5}
	fmt.Printf("&p1 = %p\n", &p1)
	fmt.Printf("&p2 = %p\n", &p2)
	fmt.Println(p1 == p2)

	// var x = &Product{}
	// var x = new(Product)
}

func Format(p Product) string {
	return fmt.Sprintf("id = %d, name = %q, cost = %0.2f", p.id, p.name, p.cost)
}

// write a function that applies the given discount (percentage) to the given product
func ApplyDiscount(p *Product, discountPercentage float64) /* do not return */ {
	p.cost = p.cost * ((100 - discountPercentage) / 100)
}
