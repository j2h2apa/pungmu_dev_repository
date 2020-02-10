package main

import (
	"fmt"
)

type Student struct {
	name  string
	class int8

	sungjuk Sungjuk
}

func (s Student) viewSungJuk() {
	fmt.Println(s.sungjuk)
}

type Sungjuk struct {
	name  string
	grade string
}

func main() {
	var s Student

	s.name = "철수"
	s.class = 1

	s.sungjuk.name = "수학"
	s.sungjuk.grade = "C"

	s.viewSungJuk()
}
