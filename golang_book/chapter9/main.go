package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}
type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Phi * c.r * c.r
}
func (c *Circle) perimeter() float64 {
	return 2 * math.Phi * c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	h := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return h * w
}

func (r *Rectangle) perimeter() float64 {
	h := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2 * (h + w)
}
func main() {
	c := Circle{x: 0, y: 0, r: 5}
	r := Rectangle{0, 0, 2, 3}
	fmt.Println("area circle 	= ", c.area())
	fmt.Println("area rectangle	= ", r.area())
	fmt.Println("perimeter circle 	= ", c.perimeter())
	fmt.Println("perimeter rectangle	= ", r.perimeter())

	fmt.Println("total area		= ", totalArea(&c, &r))
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}
func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}
