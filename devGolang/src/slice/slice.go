package main

import (
	"fmt"
)

// practice 1 : declare
func sliceDeclare() {
	var a []int

	fmt.Printf("len(a) = %d\n", len(a))
	fmt.Printf("cap(a) = %d\n", cap(a))

	b := []int{1, 2, 3, 4, 5}
	fmt.Printf("len(b) = %d\n", len(b))
	fmt.Printf("cap(b) = %d\n", cap(b))

	c := make([]int, 0, 8)
	fmt.Printf("len(c) = %d\n", len(c))
	fmt.Printf("cap(c) = %d\n", cap(c))

	c = append(c, 1)
	fmt.Printf("len(c) = %d\n", len(c))
	fmt.Printf("cap(c) = %d\n", cap(c))
}

// practice 2 : capatity 비교
// len 2 배열에서 한 개를 더 추가하면 capacity 는 배수 4가 되고 len 은 3이된다.
// 메모리를 새로 확보
func capacityCompare() {
	a := []int{1, 2}
	b := append(a, 3)

	fmt.Printf("%p %p\n", a, b)

	for i := 0; i < len(a); i++ {
		fmt.Printf("%d,", a[i])
	}
	fmt.Println()
	for i := 0; i < len(b); i++ {
		fmt.Printf("%d,", b[i])
	}
	fmt.Println()

	fmt.Printf("cap(a) / cap(b) : %d / %d\n", cap(a), cap(b))
}

// practice 3
// make 로 capacity 를 미리 확보한 상태이므로 주소는 같음
// 같은 메모리를 사용한다
func capacityCompare2() {
	a := make([]int, 2, 4)
	b := append(a, 3)

	fmt.Printf("pointer(a) / pointer(b) : %p / %p\n", a, b)

	fmt.Println(a)
	fmt.Println(b)

	b[0] = 4
	b[1] = 5

	fmt.Println(a)
	fmt.Println(b)
}

// practice 4 : 서로 다른 메모리 운영하는 방법
func independentMemory() {
	var a []int = make([]int, 2, 4)
	a[0] = 1
	a[1] = 2

	b := make([]int, len(a))

	for i := 0; i < len(a); i++ {
		b[i] = a[i]
	}

	b = append(b, 3)

	fmt.Printf("pointer(a) / pointer(b) : %p / %p\n", a, b)

	b[0] = 4
	b[1] = 5

	fmt.Println(a)
	fmt.Println(b)
}

// practice 5 : default slice
// slice array 는 원래 메모리의 정보를 pointer 한다
func defaultSlice() {
	var a []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	b := a[4:8]
	c := a[4:]
	d := a[:8]

	fmt.Printf("pointer(a) / pointer(b) : %p / %p\n", a, b)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

// practice 6 : 길이 - 1 length 예재
func removeBack(a []int) ([]int, int) {
	fmt.Printf("removeBack pointer(a) : %p\n", a)
	return a[:len(a)-1], a[len(a)-1]
}

// practice 7 : 앞에 값부터 제거하여 리턴
func removeFront(a []int) ([]int, int) {
	return a[1:], a[0]
}

func main() {
	sliceDeclare()
	capacityCompare()
	capacityCompare2()
	independentMemory()

	// slice
	defaultSlice()

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("len(a) : %d pointer(a) : %p\n", len(a), a)
	back := 0
	for i := 0; i < len(a); i++ {
		a, back = removeFront(a)
		fmt.Printf("%d, ", back)
	}
	fmt.Println(a)

	fmt.Println("============================================================")
}
