/*
Part A Declare and initialize a variable of type int with the value of 20. Display the address of and value of the variable.

Part B Declare and initialize a pointer variable of type int that points to the last variable you just created. Display the address of , value of and the value that the pointer points to.

*/

package main

import "fmt"

func main() {
	partA := 20
	fmt.Println(&partA, partA)
	partB := &partA
	fmt.Println(&partB, partB, *partB)

}
