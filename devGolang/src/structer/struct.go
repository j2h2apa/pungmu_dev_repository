package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int8
}

func (p Person) printName() {
	fmt.Print(p.name)
}

func main() {
	var p Person

	p1 := Person{"개똥이", 15}
	p2 := Person{name: "소똥이", age: 21}
	p3 := Person{name: "Carson"}
	p4 := Person{}

	fmt.Println(p, p1, p2, p3, p4)

	p.name = "Smith"
	p.age = 24

	fmt.Println(p)

	p.printName()
}
