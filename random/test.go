package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

func printCircle(c Circle) {
	fmt.Println(c.X)
	fmt.Println(c.Y)
	fmt.Println(c.Radius)
}

func main() {
	fmt.Println("hello world!")
	var c Circle
	c.X = 3
	c.Y = 6
	c.Radius = 10

	b := Circle{Point{1, 2}, 3}
	printCircle(b)
	printCircle(c)

}
