package main

import (
	"datastruct"
	"fmt"
)

func stack() {
	var stack []int = []int{}

	for i := 0; i < 5; i++ {
		stack = append(stack, i)
	}

	fmt.Println(stack)

	for len(stack) > 0 {
		var last int
		last, stack = stack[len(stack)-1], stack[:len(stack)-1]
		fmt.Println(last)
	}
}

func queue() {
	var queue []int = []int{}

	for i := 1; i < 6; i++ {
		queue = append(queue, i)
	}

	fmt.Println(queue)

	for len(queue) > 0 {
		var front int
		front, queue = queue[0], queue[1:]
		fmt.Println(front)
	}
}

func main() {
	// practice 1
	stack()
	queue()

	// practice 2
	var stack *datastruct.Stack = datastruct.NewStack()

	for i := 1; i <= 5; i++ {
		stack.Push(i)
	}

	fmt.Println("stack 2 start ==========================")

	for !stack.Empty() {
		val := stack.Pop()
		fmt.Printf("%d -> ", val)
	}

	// practice 3 : Queue
	fmt.Println("\nqueue 2 start ==========================")

	var queue *datastruct.Queue = datastruct.NewQueue()

	for i := 1; i <= 5; i++ {
		queue.Push(i)
	}

	for !queue.Empty() {
		val := queue.Pop()
		fmt.Printf("%d -> ", val)
	}

}
