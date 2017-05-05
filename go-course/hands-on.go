package main

import (
	"fmt"
	"math"
)

type square struct {
	l float64
}

type circle struct {
	r float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

func (s square) area() float64 {
	return s.l * s.l
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {

	cir := circle{5}
	sqr := square{10}
	info(cir)
	info(sqr)

}
