package main

import (
	"fmt"
)

func main() {
	for i:= 1; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
	q_2()
	q_3()

}

func q_2() {
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 {
			fmt.Println(i)
		}
	}
}

func q_3() {
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i % 3 == 0 {
			fmt.Println("Fizz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		}
	}
}



