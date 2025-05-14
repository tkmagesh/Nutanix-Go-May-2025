package main

import "fmt"

func main() {
	// product -> id, name, cost
	/*
		var p struct {
			id   int
			name string
			cost float64
		}
		p.id = 100
		p.name = "Pen"
		p.cost = 10
	*/
	var p struct {
		id   int
		name string
		cost float64
	} = struct {
		id   int
		name string
		cost float64
	}{
		id:   100,
		name: "pen",
		cost: 10,
	}
	// fmt.Println(p)
	// fmt.Printf("%#v\n", p)
	// fmt.Printf("%+v\n", p)
	fmt.Println(Format(p))
}

func Format(p struct {
	id   int
	name string
	cost float64
}) string {
	return fmt.Sprintf("id = %d, name = %q, cost = %0.2f", p.id, p.name, p.cost)
}
