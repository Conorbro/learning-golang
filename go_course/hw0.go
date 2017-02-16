// Homework 0: Hello Go!
// Due January 24, 2017 at 11:59pm
package main

import "fmt"

func main() {
	// Feel free to use the main function for testing your functions
	fmt.Println("Hello, दुनिया!")
	Fizzbuzz(20)
	fmt.Println(IsPrime(101))
	fmt.Println(IsPalindrome("ratsliveonnoevilstar"))
}

// Fizzbuzz is a classic introductory programming problem.
// If n is divisible by 3, return "Fizz"
// If n is divisible by 5, return "Buzz"
// If n is divisible by 3 and 5, return "FizzBuzz"
// Otherwise, return the empty string
func Fizzbuzz(n int) string {
	for i := 1; i <= n; i++ {
		if i%5 == 0 && i%3 == 0 {
			fmt.Println("Fizz Buzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}
	return ""
}

// IsPrime checks if the number is prime.
// You may use any prime algorithm, but you may NOT use the standard library.
func IsPrime(n int) bool {
	if n == 1 || n == 0 {
		return false
	}
	temp := n - 1
	for temp > 1 {
		if n%temp == 0 {
			return false
		}
		temp--
	}
	return true
}

// IsPalindrome checks if the string is a palindrome.
// A palindrome is a string that reads the same backward as forward.
func IsPalindrome(s string) bool {
	runes := []rune(s)
	var end = len(runes) - 1
	var temp rune
	for i := 0; i < (len(runes) / 2); i++ {
		temp = runes[end]
		runes[end] = runes[i]
		runes[i] = temp
		end--
	}
	var res = string(runes)
	if res == s {
		return true
	}
	return false
}
