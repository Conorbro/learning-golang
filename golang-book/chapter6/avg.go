package main

import "fmt"

func multi() (int, int) {
	return 5, 6
}

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func main() {
	x := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(x))
	y, z := multi()
	fmt.Println(y, z)
}
