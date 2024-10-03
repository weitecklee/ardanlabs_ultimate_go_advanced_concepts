/*
Part A: Declare a struct type to maintain information about a user (name, email and age). Create a value of this type, initialize with values and display each field.

Part B: Declare and initialize an anonymous struct type with the same three fields. Display the value.
*/

package main

import "fmt"

type user struct {
	name  string
	email string
	age   int
}

func main() {
	u1 := user{
		name:  "Dude",
		email: "dude@sweet.com",
		age:   25,
	}
	fmt.Println(u1.name, u1.email, u1.age)
	u2 := struct {
		name  string
		email string
		age   int
	}{
		name:  "Bro",
		email: "bro@cool.com",
		age:   26,
	}
	fmt.Println(u2.name, u2.email, u2.age)
}
