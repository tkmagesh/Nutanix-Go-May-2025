package main

import (
	"fmt"
	"math"
)

// ver 1.0
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// ver 2.0
type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// Refactoring
/*
func PrintArea(x interface{}) {
	// fmt.Println("Area :", x.Area())
	switch s := x.(type) {
	case Circle:
		fmt.Println("Area :", s.Area())
	case Rectangle:
		fmt.Println("Area :", s.Area())
	default:
		fmt.Println("Argument doesnot have Area() method")
	}
}
*/

/*
func PrintArea(x interface{}) {
	switch s := x.(type) {
	case interface{ Area() float64 }:
		fmt.Println("Area :", s.Area())
	default:
		fmt.Println("Argument doesnot have Area() method")
	}
}
*/

func PrintArea(x interface{ Area() float64 }) {
	fmt.Println("Area :", x.Area())
}

// ver 3.0 (introduce perimeters)
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func PrintPerimeter(x interface{ Perimeter() float64 }) {
	fmt.Println("Perimeter :", x.Perimeter())
}

// Refactoring
/*
func PrintShape(x interface {
	interface{ Area() float64 }
	interface{ Perimeter() float64 }
}) {
	PrintArea(x)      // x should be interface{ Area() float64 }
	PrintPerimeter(x) // x should be interface{ Perimeter() float64 }
}
*/

func PrintShape(x interface {
	Area() float64
	Perimeter() float64
}) {
	PrintArea(x)      // x should be interface{ Area() float64 }
	PrintPerimeter(x) // x should be interface{ Perimeter() float64 }
}

func main() {

	// ver 1.0
	c := Circle{Radius: 16}
	// fmt.Println("Area :", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintShape(c)

	// ver 2.0
	r := Rectangle{Height: 14, Width: 18}
	// fmt.Println("Area :", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintShape(r)
	// PrintArea(100)
}
