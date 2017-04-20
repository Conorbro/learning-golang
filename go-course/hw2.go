// Homework 2: Object Oriented Programming
// Due February 7, 2017 at 11:59pm
package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// Feel free to use the main function for testing your functions
	// world := struct {
	// 	English string
	// 	Spanish string
	// 	French  string
	// }{
	// 	"world",
	// 	"mundo",
	// 	"monde",
	// }
	// fmt.Printf("Hello, %s/%s/%s!", world.English, world.Spanish, world.French)

	// Q1
	// var p Price = 2595
	// fmt.Println(p.String())

	//Q2
	// RegisterItem(Prices, "cheese", 299)
	// fmt.Println(Prices)

	//Q3
	// my_items := []string{"apples", "oranges", "pears"}
	my_cart := Cart{Items: []string{"apples", "oranges", "pears", "milk"}, TotalPrice: 100}
	// fmt.Println(my_cart.hasMilk())
	// var my_cart Cart
	// my_cart.TotalPrice = 100
	// my_cart.Items = my_items

	//Q4
	// fmt.Println(my_cart.HasItem("apples"))
	// fmt.Println(my_cart.HasItem("cherries"))

	//Q5
	// my_cart.AddItem("eggs")
	// my_cart.AddItem("cheese")

	//Q6
	my_cart.Checkout()

}

// Price is the cost of something in US cents.
type Price int64

// String is the string representation of a Price
// These should be represented in US Dollars
// Example: 2595 cents => $25.95
func (p Price) String() string {
	p_float := float64(p)                                  // convert from int64 to float64
	p_float = p_float / 100                                // divide by 100 to convert from cents to dollars
	res := "$" + strconv.FormatFloat(p_float, 'f', -1, 64) // convert float64 to string
	return res
}

// Prices is a map from an item to its price.
var Prices = map[string]Price{
	"eggs":          219,
	"bread":         199,
	"milk":          295,
	"peanut butter": 445,
	"chocolate":     150,
}

// RegisterItem adds the new item in the prices map.
// If the item is already in the prices map, a warning should be displayed to the user,
// but the value should be overwritten.
// Bonus (1pt) - Use the "log" package to print the error to the user
func RegisterItem(prices map[string]Price, item string, price Price) {
	prices[item] = price
}

// Cart is a struct representing a shopping cart of items.
type Cart struct {
	Items      []string
	TotalPrice Price
}

// hasMilk returns whether the shopping cart has "milk".
func (c *Cart) hasMilk() bool {
	for _, item := range c.Items {
		if item == "milk" {
			return true
		}
	}
	return false
}

// HasItem returns whether the shopping cart has the provided item name.
func (c *Cart) HasItem(item string) bool {
	for _, i := range c.Items {
		if i == item {
			return true
		}
	}
	return false
}

// AddItem adds the provided item to the cart and update the cart balance.
// If item is not found in the prices map, then do not add it and print an error.
// Bonus (1pt) - Use the "log" package to print the error to the user
func (c *Cart) AddItem(item string) {
	c.Items = append(c.Items, item)

	if _, ok := Prices[item]; ok {
		c.TotalPrice += Prices[item]
		fmt.Println("Added", item, ":", Prices[item])
		fmt.Println(c.Items)
	} else {
		log.Printf("Item %s does not exist in prices!", item)
	}
}

// Checkout displays the final cart balance and clears the cart completely.
func (c *Cart) Checkout() {
	var totalPrice Price
	for _, item := range c.Items {
		if _, ok := Prices[item]; ok {
			totalPrice += Prices[item]
		} else {
			log.Printf("Item %s does not exist in Prices!", item)
		}
	}
	c.Items = c.Items[:0]
	fmt.Println(len(c.Items))
	fmt.Println("Total Price =", totalPrice)
}
