// Homework 1: Finger Exercises
// Due January 31, 2017 at 11:59pm
package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// Feel free to use the main function for testing your functions

	// intArr := []int{1, 2, 3, 3, 4, 5, 6, 5, 5, 2, 1}
	// sampleMap := map[string]int{
	// 	"test1": 1,
	// 	"test2": 2,
	// 	"test3": 3,
	// }

	// fmt.Println("Hello, دنيا!")
	// fmt.Println(ParsePhone("   1 --2 3 4 ---5 6 7 8 9 0----"))
	// fmt.Println(Anagram("conor broderick", "brrickode ronoc"))
	// fmt.Println(FindEvens(intArr))
	// fmt.Println(SliceProduct(intArr))
	// fmt.Println(Unique(intArr))
	// fmt.Println(InvertMap(sampleMap))
	fmt.Println(TopCharacters("abbccjdidbbhufeijaawecfo", 3))
}

// ParsePhone parses a string of numbers into the format (123) 456-7890.
// This function should handle any number of extraneous spaces and dashes.
// All inputs will have 10 numbers and maybe extra spaces and dashes.
// For example, ParsePhone("123-456-7890") => "(123) 456-7890"
//              ParsePhone("1 2 3 4 5 6 7 8 9 0") => "(123) 456-7890"
func ParsePhone(phone string) string {
	var res []rune
	res = append(res, '(')
	for _, r := range phone {
		if r < 58 && r >= 48 {
			res = append(res, r)
		}
		if len(res) == 4 {
			res = append(res, ')', ' ')
		}
		if len(res) == 9 {
			res = append(res, '-')
		}
	}
	return string(res)
}

// Anagram tests whether the two strings are anagrams of each other.
// This function is NOT case sensitive and should handle UTF-8
func Anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	s1, s2 = strings.ToLower(s1), strings.ToLower(s2)
	s1Arr, s2Arr := strings.Split(s1, ""), strings.Split(s2, "")
	sort.Strings(s1Arr)
	sort.Strings(s2Arr)
	return strings.Join(s1Arr, "") == strings.Join(s2Arr, "")
}

// FindEvens filters out all odd numbers from input slice.
// Result should retain the same ordering as the input.
func FindEvens(e []int) []int {
	var res []int
	for _, r := range e {
		if r%2 == 0 {
			res = append(res, r)
		}
	}
	return res
}

// SliceProduct returns the product of all elements in the slice.
// For example, SliceProduct([]int{1, 2, 3}) => 6
func SliceProduct(e []int) int {
	var res = 1
	for _, r := range e {
		res = res * r
	}
	return res
}

// Unique finds all distinct elements in the input array.
// Result should retain the same ordering as the input.
func Unique(e []int) []int {
	var res []int
	numberMap := make(map[int]bool)
	for _, r := range e {
		if !numberMap[r] {
			numberMap[r] = true
		}
	}
	for k := range numberMap {
		res = append(res, k)
	}
	return res
}

// InvertMap inverts a mapping of strings to ints into a mapping of ints to strings.
// Each value should become a key, and the original key will become the corresponding value.
// For this function, you can assume each value is unique.
func InvertMap(kv map[string]int) map[int]string {
	res := make(map[int]string)
	for k, v := range kv {
		res[v] = k
	}
	return res
}

// TopCharacters finds characters that appear more than k times in the string.
// The result is the set of characters along with their occurrences.
// This function MUST handle UTF-8 characters.
func TopCharacters(s string, k int) map[rune]int {
	res := make(map[rune]int)
	sSlice := []rune(s)
	for i := 0; i < len(sSlice); i++ {
		if _, ok := res[sSlice[i]]; ok {
			res[sSlice[i]]++
		} else {
			res[sSlice[i]] = 1
		}
	}
	for key, val := range res {
		if val < k {
			delete(res, key)
		}
	}
	return res
}
