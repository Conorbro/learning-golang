package main

import "fmt"

func findBig(args ...int) int {
	biggest := 0
	for _, v := range args {
		if v > biggest {
			biggest = v
		}
	}
	return biggest
}

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 7, 10}
	fmt.Println(findBig(list...))
	fmt.Println(findBig(1, 2, 3, 4, 5, 6))
}
