package main

import "math"

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base, Height float64
}

func (r Rectangle) Area() float64 { return r.Width * r.Height }

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (c Circle) Area() float64 { return math.Pi * math.Pow(c.Radius, 2) }

func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

func (t Triangle) Area() float64 { return 0.5 * t.Base * t.Height }

func (t Triangle) Perimeter() float64 { return 3 * t.Base }

func (r *Rectangle) Scale(f float64) { r.Width *= f; r.Height *= f }

func (c *Circle) Scale(f float64) { c.Radius *= f }

func (t *Triangle) Scale(f float64) { t.Base *= f; t.Height *= f }
