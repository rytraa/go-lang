package main

import (
	"fmt"
	"math"
)

// ----------------------------------------------------

type Shape interface {
	area() float64
}

// ----------------------------------------------------
type Rect struct {
	l, w float64
}

func (r *Rect) area() float64 {
	return r.l * r.w
}

// ----------------------------------------------------

type Circle struct {
	r float64
}

func (c *Circle) area() float64 {
	return c.r * c.r * math.Pi
}

// ----------------------------------------------------

func CalculateArea(x Shape) float64 {
	return x.area()
}

// ----------------------------------------------------

type MergedShapes struct {
	s []Shape
}

func (m *MergedShapes) area() (total float64) {

	for _, v := range m.s {
		total += v.area()
	}

	return

}

func main() {

	fmt.Println("Hello, World")
	fmt.Println(CalculateArea(&Circle{10}))
	fmt.Println(CalculateArea(&Rect{10, 20}))

	m := MergedShapes{s: []Shape{&Circle{5}, &Rect{6, 7}, &MergedShapes{s: []Shape{&Circle{5}, &Rect{6, 7}}}}}
	fmt.Println(CalculateArea(&m))

}
