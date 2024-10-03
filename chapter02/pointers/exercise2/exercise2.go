/*
Declare a struct type and create a value of this type. Declare a function that can change the value of some field in this struct type. Display the value before and after the call to your function.
*/

package main

import "fmt"

type person struct {
	name string
	age  int
}

func celebrateBirthday(p *person) {
	p.age++
}

func main() {
	p := person{name: "Dude", age: 25}
	fmt.Println(p.name, p.age)
	celebrateBirthday(&p)
	fmt.Println(p.name, p.age)
}
