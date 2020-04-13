package main

import (
	"fmt"
	"time"
)

func push(c chan<- int) {
	var i int
	for {
		time.Sleep(2 * time.Second)
		c <- i
		i++
	}
}

func main() {
	var c chan int = make(chan int)

	go push(c)

	for {
		select {
		case v := <-c:
			fmt.Println(v)
		default:
			fmt.Println("Idle")
			time.Sleep(1 * time.Second)
		}
	}
}
