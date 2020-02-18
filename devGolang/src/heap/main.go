package main

import (
	"datastruct"
	"fmt"
)

func main() {
	var h datastruct.Heap = datastruct.Heap{}

	// [-1, 3, -1, 45, 4], 2번째 큰 값 출력
	nums := []int{-1, 3, -1, 45, 4}

	for i := 0; i < len(nums); i++ {
		h.Push(nums[i])
		if h.Count() > 2 {
			h.Pop()
		}
	}

	fmt.Println(h.Pop())
}
