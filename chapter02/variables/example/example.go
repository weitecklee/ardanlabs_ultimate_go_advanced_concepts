package main

import "fmt"

func main() {

	// Declare variables that are set to their zero values
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("var a int \t %T [%#v]\n", a, a)
	fmt.Printf("var b string \t %T [%#v]\n", b, b)
	fmt.Printf("var c float64 \t %T [%#v]\n", c, c)
	fmt.Printf("var d bool \t %T [%#v]\n\n", d, d)

	// Declare variables and initialize
	// Using the short variable declaration operator
	aa := 10
	bb := "hello"
	cc := 3.14159
	dd := true

	fmt.Printf("var aa int \t %T [%#v]\n", aa, aa)
	fmt.Printf("var bb string \t %T [%#v]\n", bb, bb)
	fmt.Printf("var cc float64 \t %T [%#v]\n", cc, cc)
	fmt.Printf("var dd bool \t %T [%#v]\n\n", dd, dd)

	// Specify type and perform a conversion
	aaa := int32(10)
	fmt.Printf("aaa := int32(10) \t %T [%#v]\n", aaa, aaa)

}
