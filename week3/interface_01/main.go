package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Triangle struct {
	SideA float64
	SideB float64
	SideC float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func (t Triangle) Area() float64 {
	return 0.5 * t.SideA * t.SideB
}

func (t Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}

func main() {
	shapes := []Shape{
		Circle{Radius: 0.5},
		Rectangle{Width: 0.5, Height: 0.5},
		Triangle{SideA: 0.4, SideB: 0.6, SideC: 0.8},
	}
	for _, shape := range shapes {
		fmt.Printf("图形面积:%.2f,图形周长:%.2f", shape.Area(), shape.Perimeter())
	}
}
