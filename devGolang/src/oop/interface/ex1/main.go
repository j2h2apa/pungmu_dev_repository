package main

import (
	"fmt"
)

type SpoonOfJam interface {
	String() string
}

type Jam interface {
	GetOneSpoon() SpoonOfJam
}

type Bread struct {
	val string
}

func (b *Bread) PutJam(jam Jam) {
	var spoon SpoonOfJam = jam.GetOneSpoon()
	b.val += spoon.String()
}

func (b *Bread) String() string {
	return "bread " + b.val
}

type StrawberryJam struct {
}

func (j *StrawberryJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfStrawberryJam{}
}

type OrangeJam struct {
}

func (j *OrangeJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfOrangeJam{}
}

type AppleJam struct{}

func (j *AppleJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfAppleJam{}
}

type SpoonOfStrawberryJam struct {
}

type SpoonOfOrangeJam struct {
}

type SpoonOfAppleJam struct{}

func (s *SpoonOfStrawberryJam) String() string {
	return " + strawberry"
}

func (s *SpoonOfOrangeJam) String() string {
	return " + orange"
}

func (s *SpoonOfAppleJam) String() string {
	return " + apple"
}

func main() {
	fmt.Println("============ Process START =============")

	var bread *Bread = &Bread{}
	// var jam *StrawberryJam = &StrawberryJam{}
	// var jam *OrangeJam = &OrangeJam{}
	jam := &AppleJam{}

	bread.PutJam(jam)

	fmt.Println(bread)

	fmt.Println("============ Process END =============")
}
