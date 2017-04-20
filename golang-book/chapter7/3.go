package main

import "fmt"

// Add a new perimeter method to the Shape interface to calculate the perimeter of a shape. Implement the method for Circle and Rectangle.

type Shape interface {
	perimeter() float64
}

type Rect struct {
	x1, y1, x2, y2 float64
}

type Circle struct {
	x, y, r float64
}

func (r *Rect) perimeter() float64 {

}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func main() {

}
