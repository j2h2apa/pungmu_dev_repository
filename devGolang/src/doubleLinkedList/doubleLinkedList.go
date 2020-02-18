package main

import (
	"dataStruct"
	"fmt"
)

func main() {
	var list *dataStruct.LinkedList = &dataStruct.LinkedList{}

	list.AddNode(0)

	for i := 1; i < 11; i++ {
		list.AddNode(i)
	}

	list.PrintNodes()
	list.PrintReverse()

	list.RemoveNode(list.Root.Next)

	list.PrintNodes()

	list.RemoveNode(list.Root)

	list.PrintNodes()

	fmt.Printf("%d\n", list.Tail.Val)

	list.RemoveNode(list.Tail)

	list.PrintNodes()

	fmt.Printf("%d\n", list.Tail.Val)

	// slice 와의 차이 (동적 배열과의 차이)
	a := []int{1, 2, 3, 4, 5}
	a = append(a[0:2], a[3:]...)
	fmt.Println(a)

}
