package main

import (
	"fmt"
	"time"
)

func push(c chan<- int) {
	var i int

	for {
		c <- i
		i++
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var c chan int = make(chan int)

	go push(c)

	timeChan := time.After(10 * time.Second)
	tickTimerChan := time.Tick(2 * time.Second)

	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-timeChan:
			fmt.Println("timeout")
			return
		case <-tickTimerChan:
			fmt.Println("Tick")
		}
	}
}
