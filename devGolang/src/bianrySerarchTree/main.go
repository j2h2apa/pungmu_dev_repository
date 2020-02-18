package main

import (
	"datastruct"
	"fmt"
)

func main() {
	var tree *datastruct.BinaryTree = datastruct.NewBinaryTree(5)
	tree.Root.AddNode(3)
	tree.Root.AddNode(2)
	tree.Root.AddNode(4)
	tree.Root.AddNode(8)
	tree.Root.AddNode(7)
	tree.Root.AddNode(6)
	tree.Root.AddNode(10)
	tree.Root.AddNode(9)

	tree.Print()

	fmt.Println()
	if found, cnt := tree.IsValue(6); found {
		fmt.Println("found 6 cnt : ", cnt)
	} else {
		fmt.Println("Not found 6 cnt : ", cnt)
	}
}
