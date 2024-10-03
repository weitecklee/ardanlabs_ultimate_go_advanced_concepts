/*
Part A: Declare three variables that are initialized to their zero value and three declared with a literal value. Declare variables of type string, int and bool. Display the values of those variables.

Part B: Declare a new variable of type float32 and initialize the variable by converting the literal value of Pi (3.14).
*/

package main

import "fmt"

func main() {
	var stringZero string
	var intZero int
	var boolZero bool
	fmt.Printf("var stringZero string \t %T [%#v]\n", stringZero, stringZero)
	fmt.Printf("var intZero int \t %T [%#v]\n", intZero, intZero)
	fmt.Printf("var boolZero bool \t %T [%#v]\n", boolZero, boolZero)
	name := "Dude"
	age := 25
	isSweet := true
	fmt.Println(name, age, isSweet)
	pi := float32(3.14)
	fmt.Println(pi)
}
