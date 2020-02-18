package main

import (
	"fmt"
)

type Node struct {
	next *Node
	val  int
}

// practice 1 : tail loop를 이용하여 찾아 node 추가
func addNode(root *Node, val int) {
	var tail *Node

	tail = root
	for tail.next != nil {
		tail = tail.next
	}

	node := &Node{val: val}
	tail.next = node
}

// practice 2 : tail 을 알고 진행하는 방법
func addNodeKnowTail(tail *Node, val int) *Node {
	node := &Node{val: val}
	tail.next = node
	return node
}

// practice 3 : linkedlist node remove
func removeNode(node *Node, root *Node, tail *Node) (*Node, *Node) {
	if node == root {
		root = root.next
		// node 가 한개만 있을 경우 root 와 tail 둘다 nil 처리
		if root == nil {
			tail = nil
		}
		return root, tail
	}

	prev := root
	for prev.next != node {
		prev = prev.next
	}

	if node == tail {
		prev.next = nil
		tail = prev
	} else {
		prev.next = prev.next.next
	}

	return root, tail
}

func printNodes(root *Node) {
	node := root
	for node.next != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.next
	}

	fmt.Printf("%d\n", node.val)
}

func main() {
	// practice 1
	var root *Node = &Node{val: 0}

	for i := 1; i < 11; i++ {
		addNode(root, i)
	}

	node := root
	for node.next != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.next
	}

	fmt.Printf("%d\n", node.val)
	// practice 1 end

	// practice 2
	var root2 *Node = &Node{val: 0}
	var tail *Node = root2

	for i := 1; i < 11; i++ {
		tail = addNodeKnowTail(tail, i)
	}

	printNodes(root2)
	// practice 2 end

	root2, tail = removeNode(tail, root2, tail)
	printNodes(root2)
}
