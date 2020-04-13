package main

/* channel return sample */
import (
	"fmt"
)

func num(a, b int) <-chan int {
	var out chan int = make(chan int)

	go func() {
		out <- a
		out <- b
		close(out)
	}()

	return out
}

// pop return function
func sum(c <-chan int) <-chan int {
	var out chan int = make(chan int)
	go func() {
		r := 0
		for i := range c {
			r += i
		}
		out <- r
		close(out)
	}()
	return out
}

func main() {
	var c <-chan int = num(1, 2)
	var out <-chan int = sum(c)

	fmt.Println(<-out)
	if _, success := <-out; !success {
		fmt.Println("데이터 없음")
	}
}
