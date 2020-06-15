package dataStruct

import "fmt"

type TreeNode struct {
	Val    int
	Childs []*TreeNode
}

type Tree struct {
	Root *TreeNode
}

func (t *TreeNode) AddNode(n int) {
	t.Childs = append(t.Childs, &TreeNode{Val: n})
}

func (t *Tree) AddNode(n int) {
	if t.Root == nil {
		t.Root = &TreeNode{Val: n}
	} else {
		t.Root.Childs = append(t.Root.Childs, &TreeNode{Val: n})
	}
}

func (t *Tree) RecursiveDFS() {
	RecursiveDFS(t.Root)
}

// 재귀호출을 이용한 DFS 수행
func RecursiveDFS(node *TreeNode) {
	fmt.Printf("%d -> ", node.Val)

	for i := 0; i < len(node.Childs); i++ {
		RecursiveDFS(node.Childs[i])
	}
}

// Stack을 이용한 DFS 수행
func (t *Tree) StackedDFS() {
	s := []*TreeNode{}
	s = append(s, t.Root)

	for len(s) > 0 {
		var last *TreeNode
		last, s = s[len(s)-1], s[:len(s)-1]

		fmt.Printf("%d -> ", last.Val)

		for i := 0; i < len(last.Childs); i++ {
			//for i := len(last.Childs) - 1; i >= 0; i-- {
			s = append(s, last.Childs[i])
		}
	}
}

// Queue를 이용한 BFS 수행
func (t *Tree) QueuedBFS() {
	q := []*TreeNode{}
	q = append(q, t.Root)

	for len(q) > 0 {
		var first *TreeNode
		first, q = q[0], q[1:]

		fmt.Printf("%d -> ", first.Val)

		for i := 0; i < len(first.Childs); i++ {
			q = append(q, first.Childs[i])
		}
	}
}
