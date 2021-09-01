package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func distance(x1 float64, x2 float64, y1 float64, y2 float64) float64 {
	return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func main() {

	fmt.Println("Hello, World")
	var c Circle
	fmt.Println(c)

	x := Circle{}
	fmt.Println(x)

	y := Circle{x: 50, y: 60, r: 2}
	fmt.Println(y)

	z := new(Circle) // Heap Memory
	z.x = 50
	fmt.Println(*z)

	r := &Circle{} // Stack Memory
	fmt.Println(r)

	fmt.Println(circleArea(&y))
	fmt.Println(y.area())

	a := Rectangle{2.0, 3.0, 5.0, 6.0}
	fmt.Println(a.area())

}
