package dataStruct

import "fmt"

type BinaryTreeNode struct {
	Val   int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

type BinaryTree struct {
	Root *BinaryTreeNode
}

func NewBinaryTree(v int) *BinaryTree {
	tree := &BinaryTree{}
	tree.Root = &BinaryTreeNode{Val: v}

	return tree
}

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

type DepthNode struct {
	depth int
	node  *BinaryTreeNode
}

func (t *BinaryTree) DispBinaryTree() {
	q := []DepthNode{}
	q = append(q, DepthNode{depth: 0, node: t.Root})
	curDepth := 0

	for len(q) > 0 {
		var f DepthNode
		f, q = q[0], q[1:]

		if f.depth != curDepth {
			fmt.Println()
			curDepth = f.depth
		}

		fmt.Print(f.node.Val, " ")

		if f.node.left != nil {
			q = append(q, DepthNode{depth: curDepth + 1, node: f.node.left})
		}

		if f.node.right != nil {
			q = append(q, DepthNode{depth: curDepth + 1, node: f.node.right})
		}
	}
}

func (t *BinaryTree) SearchBinaryTree(v int) (bool, int) {
	return t.Root.SearchBinaryTree(v, 1)
}

func (n *BinaryTreeNode) SearchBinaryTree(v int, cnt int) (bool, int) {
	if n.Val == v {
		return true, cnt
	} else if n.Val > v {
		if n.left != nil {
			return n.left.SearchBinaryTree(v, cnt+1)
		}
		return false, cnt
	} else {
		if n.right != nil {
			return n.right.SearchBinaryTree(v, cnt+1)
		}
		return false, cnt
	}
}
