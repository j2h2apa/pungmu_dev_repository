package main

import (
	"fmt"
)

// 영문 한글 byte array 일 경우 예제 (한글이 왜 깨질까?)
func stringArray() {
	var s string = "Hello 병우"

	fmt.Println("len(s)", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Print(s[i], ",")
	}

	fmt.Println()

	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[i]), ",")
	}

	fmt.Println()
}

// golang 의 rune 과 byte array 의 관계
func stringArrayToRuneArray() {
	var s string = "Hello 월드"

	s2 := []rune(s)
	fmt.Println("lens2 : ", len(s2))
	for i := 0; i < len(s2); i++ {
		fmt.Print(string(s2[i]), ", ")
	}

	fmt.Println()
}

// byte array(string) copy
func arrayClone() {
	// practice 1
	arr := [5]int{1, 2, 3, 4, 5}
	clone := [5]int{}

	for i := 0; i < len(arr); i++ {
		clone[i] = arr[i]
	}

	fmt.Println(clone)
}

// byte array reverse clone
func reverseClone() {
	// practice 1 복사를 이용
	arr := [5]int{1, 2, 3, 4, 5}
	var temp [5]int

	for i := 0; i < len(arr); i++ {
		temp[i] = arr[len(arr)-1-i]
	}

	fmt.Println("practice1 reverse clone : ", temp)

	// practice 2 위치 변경
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}

	fmt.Println("practice2 reverse clone : ", arr)
}

// redix sort
// 범위가 정해져 있어야 한다. 갯수가 작아야 한다
func redixSort() {
	arr := [11]int{0, 5, 4, 9, 1, 2, 8, 3, 6, 4, 5}
	// 0 ~ 9 사이 숫자 카운트 이므로 index 는 10
	var temp [10]int

	for i := 0; i < len(arr); i++ {
		idx := arr[i]
		temp[idx]++
	}

	idx := 0
	for i := 0; i < len(temp); i++ {
		for j := 0; j < temp[i]; j++ {
			arr[idx] = i
			idx++
		}
	}

	fmt.Println(arr)
}

func main() {
	var A [10]int

	for i := 0; i < len(A); i++ {
		A[i] = i * i
	}

	fmt.Println(A)
	stringArray()
	stringArrayToRuneArray()
	arrayClone()
	reverseClone()
	redixSort()
}
