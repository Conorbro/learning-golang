package main

// Homework 3: Interfaces
// Due February 14, 2017 at 11:59pm

import (
	"fmt"
	"sort"
)

var peopleCount = 0

func main() {
	// Q1
	p1 := NewPerson("aaa", "aaa")
	p2 := NewPerson("aaa", "aaa")
	p3 := NewPerson("aaa", "aaa")
	// p2 := NewPerson("Aisling", "Mangan")
	// p3 := NewPerson("Bory", "Mangan-Broderick")
	var people PersonSlice
	people = []Person{*p1, *p2, *p3}
	// sort.Sort(people)
	// fmt.Println(people)

	// Q2
	fmt.Println(IsPalindrome(people))

}

// Problem 1: Sorting Names
// Sorting in Go is done through interfaces!
// To sort a collection (such as a slice), the type must satisfy sort.Interface,
// which requires 3 methods: Len() int, Less(i, j int) bool, and Swap(i, j int).
// To actually sort a slice, you need to first implement all 3 methods on a
// custom type, and then call sort.Sort on your the PersonSlice type.
// See the Go documentation: https://golang.org/pkg/sort/ for full details.

// Person stores a simple profile. These should be sorted by alphabetical order
// by last name, followed by the first name, followed by the ID. You can assume
// the ID will be unique, but the names need not be unique.
// Sorting should be case-sensitive and UTF-8 aware.
type Person struct {
	ID        int
	FirstName string
	LastName  string
}

// PersonSlice is a slice of people
type PersonSlice []Person

// Len returns length of person slice
func (p PersonSlice) Len() int {
	return len(p)
}

// Less returns true if i is less than j and false otherwise
func (p PersonSlice) Less(i, j int) bool {
	if p[i].FirstName == p[j].FirstName && p[i].LastName == p[j].LastName {
		return p[i].ID < p[j].ID
	} else if p[i].LastName == p[j].LastName {
		return p[i].FirstName < p[j].FirstName
	}
	return p[i].LastName < p[j].LastName
}

// Swap swaps person of id i's position with person of id j's position in the PersonSlice
func (p PersonSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// NewPerson is a constructor for Person. ID should be assigned automatically in
// sequential order, starting at 1 for the first Person created.
func NewPerson(first, last string) *Person {
	p := new(Person)
	p.FirstName = first
	p.LastName = last
	// peopleCount++
	p.ID = peopleCount
	return p
}

// Problem 2: IsPalindrome Redux
// Using a function that simply requires sort.Interface, you should be able to
// check if a sequence is a palindrome. You may use, adapt, or modify your code
// from HW0. Note that the input does not need to be a string: any type which
// satisfies sort.Interface can (and will) be used to test. This means that the
// only functionality you should use should come from the sort.Interface methods
// Ex: [1, 2, 1] => true

// IsPalindrome checks if the string is a palindrome.
// A palindrome is a string that reads the same backward as forward.
func IsPalindrome(s sort.Interface) bool {
	end := s.Len() - 1
	strt := 0
	for i := 0; i < s.Len()/2; i++ {
		if s.Less(strt, end) {
			return false
		}
		strt++
		end--
	}
	return true
}

// Problem 3: Functional Programming
// Write a function Fold which applies a function repeatedly on a slice,
// producing a single value via repeated application of an input function.
// The behavior of Fold should be as follows:
//   - When s is empty, return v (default value)
//   - When s has 1 value (x0), apply f once: f(v, x0)
//   - When s has 2 values (x0, x1), apply f twice, from left to right: f(f(v, x0), x1)
//   - Continue this pattern recursively to obtain the final result.

// Fold applies a left to right application of f on s starting with v.
// Note the argument signature of f - func(int, int) int.
// This means f is a function which has 2 int arguments and returns an int.
func Fold(s []int, v int, f func(int, int) int) int {

	if len(s) > 1 {
		return f(v, s[0])
	}

	return v
}
