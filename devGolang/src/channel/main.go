package main

import (
	"fmt"
	"runtime"
	"time"
)

func pop(c *chan int) {
	fmt.Println("pop func")
	// channel 에 값이 들어올 때 까지 대기
	v := <-(*c) // pop
	fmt.Println(v)
}

func sum(a, b int, c chan int) {
	c <- a + b
}

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
	// case 1
	var c chan int
	// 2번째 인자가 0일 경우 해당 value pop 되지 않으면 dead lock
	c = make(chan int, 0)

	go pop(&c)

	c <- 10 // push

	fmt.Println("End of program")

	// case 2 carfactory sample
	chan1 := make(chan Car)
	chan2 := make(chan Car)
	chan3 := make(chan Car)

	planeChan1 := make(chan Plane)
	planeChan2 := make(chan Plane)
	planeChan3 := make(chan Plane)

	go StartWork(chan1)
	go StartPlaneWork(planeChan1)
	go MakeTire(chan1, chan2, planeChan1, planeChan2)
	go MakeEngine(chan2, chan3, planeChan2, planeChan3)

	for {
		select {
		case result := <-chan3:
			fmt.Println(result.val)
		case result := <-planeChan3:
			fmt.Println(result.val)
		}
	}

	// case 3 inner function
	fmt.Println("================= case 3 start =================")
	var ch chan int = make(chan int)

	go func() {
		ch <- 123 // 123 push
	}()

	var i int
	i = <-ch
	fmt.Println(i)

	// case 4 go thread 를 이용하여 우 정수 값 더하기
	fmt.Println("================= case 4 start =================")
	go sum(100, 200, ch)
	i = <-ch
	fmt.Println(i)

	// case 5 go rutine 의 종료 bool channel 을 이용하여 waiting pattern
	fmt.Println("================= case 5 start =================")
	var done chan bool = make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		done <- true
	}()

	// waitng until go thread end
	<-done
	// close 하지 않은 상태에서 체크 시 deadlock
	close(done)
	if _, success := <-done; !success {
		fmt.Println("done channel 데이터 없음")
	}

	// case 5 go / main thread switching
	fmt.Println("================= case 6 start =================")
	done2 := make(chan bool)
	var count int = 3

	go func() {
		for i := 0; i < count; i++ {
			done2 <- true // push sync channel
			fmt.Println("go thread :", i)
			time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i < count; i++ {
		<-done2 // pop sync channel
		fmt.Println("main thread :", i)
	}

	close(done2)
	if _, success := <-done2; !success {
		fmt.Println("done2 channel 데이터 없음")
	}

	// case 6 buffed channel
	// 사용할 logical cpu count 설정
	fmt.Println("================= case 7 start =================")
	runtime.GOMAXPROCS(runtime.NumCPU())

	var done3 chan bool = make(chan bool, 2)
	count = 4
	go func() {
		for i := 0; i < count; i++ {
			done3 <- true // channel push
			fmt.Println("고루틴 :", i)
		}
	}()

	for i := 0; i < count; i++ {
		<-done3 // channel pop
		fmt.Println("메인 함수 :", i)
	}

	close(done3)
	if _, success := <-done3; !success {
		fmt.Println("done3 데이터 없음")
	}

	// case 7 channel parameter 각각 송수신 전용
	// fmt.Println("================= case 8 start =================")

	// var chan4 chan int = make(chan int)

	// go producer(chan4)
	// go consumer(chan4)

	// fmt.Scanln()

	// close(chan4)
	// if _, success := <-chan4; !success {
	// 	fmt.Println("chan4 데이터 없음")
	// }
}
