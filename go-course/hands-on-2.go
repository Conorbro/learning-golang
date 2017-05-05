package main

import "fmt"

type person struct {
	fname string
	lname string
	age   int
}

type secretAgent struct {
	person
}

func (p person) speak() {
	fmt.Println("Hello!")
}

func (sa secretAgent) speak() {
	fmt.Println("Hello Bond!")
}

func makeSpeak(h human) {
	h.speak()
}

type human interface {
	speak()
}

func main() {
	p := person{"Conor", "Broderick", 24}
	sa := secretAgent{person{"James", "Bond", 33}}
	fmt.Println(p.fname)
	p.speak()
	fmt.Println(sa.age)
	sa.speak()
	sa.person.speak()

	makeSpeak(p)
	makeSpeak(sa)

}
