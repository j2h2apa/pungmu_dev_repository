package main

import (
	"fmt"
)

type Student struct {
	name  string
	age   int
	grade int
}

func (t *Student) SetName(newName string) {
	t.name = newName
}

func main() {
	a := Student{"aaa", 20, 10}

	a.SetName("유병우")

	fmt.Println(a)
}
