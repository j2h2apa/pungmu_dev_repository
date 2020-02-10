package main

import (
	"fmt"
)

type Student struct {
	name string
	age  int

	grade string
	class string
}

func (s *Student) printSungjuk() {
	fmt.Println(s.class, s.grade)
}

func (s *Student) inputSungjuk(class, grade string) {
	s.class = class
	s.grade = grade
}

// practice 1
func defaultPointer() {
	var p *int
	var a int
	var b int

	p = &a
	a = 3
	b = 2

	fmt.Println(a)
	fmt.Println(p)
	fmt.Println(*p)

	p = &b

	fmt.Println(b)
	fmt.Println(p)
	fmt.Println(*p)
}

// practice 2 : pointer 존재 이유
func reasonPointer() {
	var a int
	a = 1

	increase(a)

	fmt.Println(a)

	increasePointer(&a)

	fmt.Println(a)
}

func increase(x int) {
	x++
}

func increasePointer(x *int) {
	*x++
}

// pratice 3 struct pointer
func pointerStruct() {
	var s Student = Student{name: "byungwoo", age: 23, class: "수학", grade: "A+"}

	s.inputSungjuk("과학", "C")
	s.printSungjuk()
}

func main() {
	fmt.Println(">>>>>>>>>>>>>>>> pointer <<<<<<<<<<<<<<<")
	defaultPointer()
	reasonPointer()
	pointerStruct()
}
