package main

import (
	"fmt"
	"time"
)

/*
	select {
	case <-chan1:
		chan1 에 값이 들어왔을 때 실행
	case <-chan2:
		chan2 에 값이 들어왔을 때 실행
	case <-chan3:
		chan3 에 값이 들어왔을 때 실행
	default:
		defaule 가 없을 경우 각 case 절에서 channel 대기를 하나 default 가 있을 경우
		모두 대기 상태면 실행된다
	}
*/

func main() {
	var c1 chan int = make(chan int)
	var c2 chan string = make(chan string)

	go func() {
		for {
			c1 <- 10
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for {
			c2 <- "hello, j2h2s2apa"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			case i := <-c1:
				fmt.Println("c1 :", i)
			case s := <-c2:
				fmt.Println("c2 :", s)
			}
		}
	}()

	time.Sleep(10 * time.Second)
}
