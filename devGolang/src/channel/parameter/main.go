package main

import (
	"fmt"
)

// push 전용 channel parameter
func producer(c chan<- int) {
	for i := 0; i < 5; i++ {
		c <- i
		fmt.Println("c chan <- :", i)
	}
	c <- 100

	// fmt.Println(<-c) // 같은 go thread 안에서 channel pop 시 컴파일 오류
}

// pop 전용 channel parameter
func consumer(c <-chan int) {
	for i := range c {
		fmt.Println("c <- chan int :", i)
	}

	fmt.Println(<-c)

	// c <- 1 // 같은 go thread 안에서 channel push 에러
}

func main() {
	var chan4 chan int = make(chan int)

	go producer(chan4)
	go consumer(chan4)

	fmt.Scanln()
}
