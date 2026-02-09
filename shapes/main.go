package main

import (
	"fmt"
)

func main() {
	r := Rectangle{Width: 3, Height: 4}
	c := Circle{Radius: 5}
	t := Triangle{Base: 6, Height: 7}

	fmt.Printf("Rectangle Area: %.2f, Perimeter: %.2f\n", r.Area(), r.Perimeter())
	fmt.Printf("Circle Area: %.2f, Perimeter: %.2f\n", c.Area(), c.Perimeter())
	fmt.Printf("Triangle Area: %.2f, Perimeter: %.2f\n", t.Area(), t.Perimeter())

	r.Scale(2)
	c.Scale(2)
	t.Scale(2)

	fmt.Printf("Scaled Rectangle Area: %.2f, Perimeter: %.2f\n", r.Area(), r.Perimeter())
	fmt.Printf("Scaled Circle Area: %.2f, Perimeter: %.2f\n", c.Area(), c.Perimeter())
	fmt.Printf("Scaled Triangle Area: %.2f, Perimeter: %.2f\n", t.Area(), t.Perimeter())
}
