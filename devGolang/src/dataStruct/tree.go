package datastruct

import (
	"fmt"
)

/*TreeNode : Tree 자료구조의 각 node*/
type TreeNode struct {
	Val    int
	Childs []*TreeNode
}

/*Tree : TreeNode 의 root 운영*/
type Tree struct {
	Root *TreeNode
}

/*AddNode : Node 추가*/
func (t *Tree) AddNode(val int) {
	// root 가 없다면
	if t.Root == nil {
		t.Root = &TreeNode{Val: val}
	} else {
		t.Root.Childs = append(t.Root.Childs, &TreeNode{Val: val})
	}
}

/*AddNode : TreeNode 에 child node 추가*/
func (t *TreeNode) AddNode(val int) {
	t.Childs = append(t.Childs, &TreeNode{Val: val})
}

/*DFS1 : tree struct method*/
func (t *Tree) DFS1() {
	DFS1(t.Root)
}

/*DFS1 : recursive call 재귀를 활용한 순환*/
func DFS1(node *TreeNode) {
	fmt.Printf("%d->", node.Val)

	for i := 0; i < len(node.Childs); i++ {
		DFS1(node.Childs[i])
	}
}

/*DFS2 : stack(slice) 활용한 순환*/
func (t *Tree) DFS2() {
	s := []*TreeNode{}
	s = append(s, t.Root)

	for len(s) > 0 {
		var last *TreeNode
		// stack pop
		last, s = s[len(s)-1], s[:len(s)-1]

		fmt.Printf("%d->", last.Val)

		for i := 0; i < len(last.Childs); i++ {
			s = append(s, last.Childs[i])
		}
	}
}

/*BFS : queue 를 활용한 sibiling 탐색*/
func (t *Tree) BFS() {
	var queue []*TreeNode = []*TreeNode{}
	queue = append(queue, t.Root)

	for len(queue) > 0 {
		var first *TreeNode
		// FIFO, queue reslicing
		first, queue = queue[0], queue[1:]

		fmt.Printf("%d->", first.Val)

		for i := 0; i < len(first.Childs); i++ {
			queue = append(queue, first.Childs[i])
		}
	}
}
