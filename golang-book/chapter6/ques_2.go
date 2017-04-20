package main

import "fmt"

func f(i int) (int, bool) {
	if i%2 == 0 {
		return i / 2, true
	}
	return i / 2, false
}

func main() {
	fmt.Println(f(2))
	fmt.Println(f(3))
}
