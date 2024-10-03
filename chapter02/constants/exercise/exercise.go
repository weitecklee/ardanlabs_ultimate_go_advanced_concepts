package main

import "fmt"

const (
	one         = 1
	two float64 = 2
)

func main() {
	fmt.Println(one, two)
	var half float64 = one / two
	fmt.Println(half)
}
