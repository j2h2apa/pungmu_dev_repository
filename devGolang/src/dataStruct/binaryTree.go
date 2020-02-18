package datastruct

// 최소신장트리를 만들어야 한다 (AVL 트리)

import (
	"fmt"
)

/*BinaryTreeNode : */
type BinaryTreeNode struct {
	left  *BinaryTreeNode
	right *BinaryTreeNode
	Val   int
}

/*BinaryTree : root node 가진 struct*/
type BinaryTree struct {
	Root *BinaryTreeNode
}

/*NewBinaryTree : root 를 포함한 BinaryTree 최초 생성*/
func NewBinaryTree(v int) *BinaryTree {
	var tree *BinaryTree = &BinaryTree{}
	tree.Root = &BinaryTreeNode{Val: v}
	return tree
}

/*AddNode : node 추가*/
func (n *BinaryTreeNode) AddNode(v int) *BinaryTreeNode {
	if n.Val > v {
		if n.left == nil {
			n.left = &BinaryTreeNode{Val: v}
			return n.left
		} else {
			return n.left.AddNode(v)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryTreeNode{Val: v}
			return n.right
		} else {
			return n.right.AddNode(v)
		}
	}
}

type depthNode struct {
	depth int
	node  *BinaryTreeNode
}

/*Print : BST node print (BFS use queue)*/
func (t *BinaryTree) Print() {
	var queue []depthNode = []depthNode{}
	queue = append(queue, depthNode{depth: 0, node: t.Root})
	currentDepth := 0

	for len(queue) > 0 {
		var first depthNode
		// queue pop
		first, queue = queue[0], queue[1:]

		if first.depth != currentDepth {
			fmt.Println()
			currentDepth = first.depth
		}

		fmt.Print(first.node.Val, " ")

		/* 최하위 node 는 정보가 없어 패스 (BFS) */
		if first.node.left != nil {
			queue = append(queue, depthNode{depth: currentDepth + 1, node: first.node.left})
		}

		if first.node.right != nil {
			queue = append(queue, depthNode{depth: currentDepth + 1, node: first.node.right})
		}
	}
}

/*IsValue : value search return bool*/
func (t *BinaryTree) IsValue(v int) (bool, int) {
	return t.Root.isValueNode(v, 1)
}

/*searchNode : BinaryTreeNode 에서 value 를 검색하여 bool 리턴*/
func (n *BinaryTreeNode) isValueNode(v int, cnt int) (bool, int) {
	if n.Val == v {
		return true, cnt
	}

	if n.Val > v {
		if n.left != nil {
			// left 계속 검색
			return n.left.isValueNode(v, cnt+1)
		}
		return false, cnt
	} else {
		if n.right != nil {
			// right 계속 검색
			return n.right.isValueNode(v, cnt+1)
		}
		return false, cnt
	}
}
