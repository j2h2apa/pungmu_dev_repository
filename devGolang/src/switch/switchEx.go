package main

import "fmt"

func main() {
	x := 33

	// switch (값) 값이 없을 경우 true 로 인식한다. case 조건이 true 인 경우 처음
	//         조건에 맞는 case 만 실행된다.
	switch 30 {
	case x - 1:
		fmt.Printf("x = %d", x-1)
	case x - 2:
		fmt.Printf("x = %d", x-2)
	case x - 3:
		fmt.Printf("x = %d", x-3)
	default:
		fmt.Println("해당하는 숫자가 없음.")
	}
}
