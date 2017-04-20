package main

import "fmt"

func makeOddGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 3
		return ret
	}
}

func main() {
	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
}
