package main

/* gorutine exit lable */

import (
	"time"
)

func run1(done chan bool) {
	time.Sleep(1 * time.Second)
	done <- true
}

func run2(done chan bool) {
	time.Sleep(2 * time.Second)
	done <- true
}

func main() {
	var done1 chan bool = make(chan bool)
	done2 := make(chan bool)

	go run1(done1)
	go run2(done2)
	/* EXIT 레이블 후 다음 구문으로 점프 프로그램 종료 */
EXIT:
	for {
		select {
		case <-done1:
			println("run1 완료")
		case <-done2:
			println("run2 완료")
			break EXIT
		}
	}
}
