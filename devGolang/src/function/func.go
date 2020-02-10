package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d + %d = %d\n", i, i+2, add(i, i+2))
	}

	// practice 1
	var a, b int = func1(2, 3)
	fmt.Println(a, b)

	// practice 2
	f1(10)

	// practice 3 recursive call
	var s int = sum(10, 0)
	fmt.Printf("practice 3 value of recursive call : %d\n", s)

	// practice 4 fibonacci
	s = fibonacci(10)
	fmt.Printf("practice 4 fibonacci : %d\n", s)
}

// practice 1
func add(x int, y int) int {
	return x + y
}

func func1(x, y int) (int, int) {
	func2(x, y)
	return y, x
}

func func2(x, y int) {
	fmt.Println("func2 ", x, y)
}

// practice 2 (recursive call)
func f1(x int) {
	if x == 0 {
		fmt.Println("recursive call end!!")
		return
	}

	fmt.Println(x)
	f1(x - 1)
}

// practice 3
func sum(x, s int) int {
	if x == 0 {
		return s
	}
	s += x
	return sum(x-1, s)
}

// practice 4 fibonacci
func fibonacci(x int) int {
	if x == 0 {
		return 1
	}
	if x == 1 {
		return 1
	}

	return fibonacci(x-1) + fibonacci(x-2)
}
