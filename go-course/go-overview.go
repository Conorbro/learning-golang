package main

import "fmt"

type gator int

type flamingo bool

type swampCreature interface {
	greeting() string
}

func bayou(s swampCreature) {
	fmt.Println(s.greeting())
}

type Person struct {
	fName   string
	lName   string
	favFood []string
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

type transportation interface {
	transportationDevice() string
}

func report(t transportation) {
	fmt.Println(t.transportationDevice())
}

func (t truck) transportationDevice() string {
	return "trucks do things"
}

func (s sedan) transportationDevice() string {
	return "sedans do things"
}

func (p Person) walk() string {
	return fmt.Sprintln(p.fName, " is walking")
}

func (g gator) greeting() string {
	return "Hello, I am a gator"
}

func (f flamingo) greeting() string {
	return "Hello, I am a flamingo"
}

func main() {

	var g1 gator
	g1 = 10
	g1.greeting()
	fmt.Println(g1)
	fmt.Printf("%T\n", g1)

	bayou(g1)
	var x int
	x = 12
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = int(g1)
	fmt.Println(x)

	p1 := Person{"Conor", "Broderick", []string{"apple", "orange", "toast"}}
	fmt.Println(p1.fName)
	fmt.Println(p1.walk())
	fmt.Println(p1.favFood)

	for _, v := range p1.favFood {
		fmt.Println(v)
	}

	slice := []int{1, 2, 3, 4, 5, 6}

	testMap := map[string]int{
		"test":  1,
		"hello": 2,
	}

	fmt.Println(slice)

	for i, _ := range slice {
		fmt.Println(i)
	}

	for i, v := range slice {
		fmt.Println(i, v)
	}

	// fmt.Println(testMap)

	for k, v := range testMap {
		fmt.Println(k, v)
	}

	for k, _ := range testMap {
		fmt.Println(k)
	}

	tr := truck{
		vehicle{4, "blue"},
		true,
	}
	sd := sedan{
		vehicle{2, "green"},
		false,
	}

	fmt.Println(tr)
	fmt.Println(sd)
	fmt.Println(tr.color)
	fmt.Println(sd.color)
	fmt.Println(tr.transportationDevice())
	fmt.Println(sd.transportationDevice())
	report(tr)
	report(sd)

	s := "Im sorry dave i can't let you do that"
	fmt.Println(s)
	fmt.Println([]byte(s))
	fmt.Println(string([]byte(s)))
	fmt.Println(string([]byte(s)[:14]))
	fmt.Println(string([]byte(s)[10:22]))
	fmt.Println(string([]byte(s)[17:]))

	for _, v := range []byte(s) {
		fmt.Println(string(v))
	}

}
