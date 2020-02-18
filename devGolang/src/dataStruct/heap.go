package datastruct

import (
	"fmt"
)

/*Heap :*/
type Heap struct {
	list []int
}

/*Push : Heap Push*/
func (h *Heap) Push(v int) {
	h.list = append(h.list, v)

	idx := len(h.list) - 1
	for idx >= 0 {
		parentIdx := (idx - 1) / 2
		// parent 가 없을 경우
		if parentIdx < 0 {
			break
		}
		/*
			minHeap : if h.list[idx] < h.list[parentIdx]
			maxHeap : if h.list[idx] > h.list[parentIdx]
		*/
		if h.list[idx] < h.list[parentIdx] {
			h.list[idx], h.list[parentIdx] = h.list[parentIdx], h.list[idx]
			idx = parentIdx
		} else {
			break
		}
	}
}

/*Print :*/
func (h *Heap) Print() {
	fmt.Println(h.list)
}

/*Pop : Heap Pop*/
func (h *Heap) Pop() int {
	if len(h.list) == 0 {
		return 0
	}

	var top int = h.list[0]
	var last int = h.list[len(h.list)-1]
	// last node remove
	h.list = h.list[:len(h.list)-1]

	// child node 가 없다면 리턴
	if len(h.list) == 0 {
		return top
	}

	// last node move at top
	h.list[0] = last

	idx := 0
	for idx < len(h.list) {
		var leftIdx, rightIdx, swapIdx int = idx*2 + 1, idx*2 + 2, -1

		// leftIdx 가 전체 길이보다 크면 더 이상 node 가 없다는 의미
		// left node index 가 더 크면 당연히 구조상 right node 는 index 가 더 크므로 존재할 수 없다
		if leftIdx >= len(h.list) {
			break
		}
		/*
			maxHeap : h.list[leftIdx] > h.list[idx]
			minHeap : h.list[leftIdx] < h.list[idx]
		*/
		if h.list[leftIdx] < h.list[idx] {
			swapIdx = leftIdx
		}

		// right node index 가 slice length 보다 작다면
		if rightIdx < len(h.list) {
			/*
				maxHeap : h.list[rightIdx] > h.list[idx]
				minHeap : h.list[rightIdx] < h.list[idx]
			*/
			if h.list[rightIdx] < h.list[idx] {
				// last 보다 큰 left node가 존재하고 left node 보다 right node 가 더 크다면
				/*
					maxHeap : h.list[swapIdx] < h.list[rightIdx]
					minHeap : h.list[swapIdx] > h.list[rightIdx]
				*/
				if swapIdx < 0 || (h.list[swapIdx] > h.list[rightIdx]) {
					swapIdx = rightIdx
				}
			}
		}

		if swapIdx < 0 {
			break
		}

		h.list[idx], h.list[swapIdx] = h.list[swapIdx], h.list[idx]
		idx = swapIdx
	}

	return top
}

/*Count : Heap 의 length count return */
func (h *Heap) Count() int {
	return len(h.list)
}
