package controllers

import (
	"fmt"

	dsPkg "lec.youtube.tucker/quiz01/dataStruct"
)

func DispLinkedList() {
	list := &dsPkg.LinkedList{}

	list.AddNode(0) // root node를 추가한다.

	for i := 1; i < 10; i++ {
		list.AddNode(i) // linked-list에 node를 추가한다.
	}

	list.PrintNodes()
	list.RemoveNode(list.Root.Next) // root 다음 node를 삭제한다.
	list.PrintNodes()
	list.RemoveNode(list.Root) // root node를 삭제한다.
	list.PrintNodes()
	list.RemoveNode(list.Tail) // tail node를 삭제한다.
	list.PrintNodes()

	fmt.Printf("tail: %d\n", list.Tail.Val)

	list.PrintReverseNodes()
}

// Stack with Slice - FILO
func DispSliceStack() {
	fmt.Println("------- DispSliceStack() --------")
	stack := []int{}
	fmt.Print("IN:\t")
	for i := 1; i <= 5; i++ {
		stack = append(stack, i)
		fmt.Printf("%d -> ", i)
	}

	fmt.Print("\nOUT:\t")
	for len(stack) > 0 {
		var last int
		last, stack = stack[len(stack)-1], stack[:len(stack)-1]
		fmt.Printf("%d -> ", last)
	}
	fmt.Println("\n---------------------------------")
}

// Queue with Slice - FIFO
func DispSliceQueue() {
	fmt.Println("------- DispSliceQueue() --------")
	queue := []int{}
	fmt.Print("IN:\t")
	for i := 1; i <= 5; i++ {
		queue = append(queue, i)
		fmt.Printf("%d -> ", i)
	}

	fmt.Print("\nOUT:\t")
	for len(queue) > 0 {
		var front int
		front, queue = queue[0], queue[1:]
		fmt.Printf("%d -> ", front)
	}
	fmt.Println("\n---------------------------------")
}

// Stack with LinkedList - FILO
func DispLinkedListStack() {
	fmt.Println("----- DispLinkedListStack() -----")
	stack := dsPkg.NewStack()
	fmt.Print("IN:\t")
	for i := 1; i <= 5; i++ {
		stack.Push(i)
		fmt.Printf("%d -> ", i)
	}

	fmt.Print("\nOUT:\t")
	for !stack.IsEmpty() {
		val := stack.Pop()
		fmt.Printf("%d -> ", val)
	}
	fmt.Println("\n---------------------------------")
}

// Queue with LinkedList - FIFO
func DispLinkedListQueue() {
	fmt.Println("----- DispLinkedListQueue() -----")
	queue := dsPkg.NewQueue()
	fmt.Print("IN:\t")
	for i := 1; i <= 5; i++ {
		queue.Push(i)
		fmt.Printf("%d -> ", i)
	}

	fmt.Print("\nOUT:\t")
	for !queue.IsEmpty() {
		val := queue.Pop()
		fmt.Printf("%d -> ", val)
	}
	fmt.Println("\n---------------------------------")
}

// 재귀호출을 이용한 TreeNode 출력(DFS[깊이우선]방식)
func DispRecursivedDFS() {
	fmt.Println("-------------- DispRecursivedDFS() ---------------")
	tree := dsPkg.Tree{}
	n := 1
	tree.AddNode(n)
	n++

	for i := 0; i < 3; i++ {
		tree.Root.AddNode(n)
		n++
	}

	for i := 0; i < len(tree.Root.Childs); i++ {
		for j := 0; j < 2; j++ {
			tree.Root.Childs[i].AddNode(n)
			n++
		}
	}

	tree.RecursiveDFS()
	fmt.Println("\n--------------------------------------------------")
}

// Stack을 이용한 TreeNode 출력(DFS[깊이우선]방식)
func DispStackedDFS() {
	fmt.Println("--------------- DispStackedDFS() -----------------")
	tree := dsPkg.Tree{}
	n := 1
	tree.AddNode(n)
	n++

	for i := 0; i < 3; i++ {
		tree.Root.AddNode(n)
		n++
	}

	for i := 0; i < len(tree.Root.Childs); i++ {
		for j := 0; j < 2; j++ {
			tree.Root.Childs[i].AddNode(n)
			n++
		}
	}

	tree.StackedDFS()
	fmt.Println("\n--------------------------------------------------")
}

// Queue를 이용한 TreeNode 출력(BFS[넓이우선]방식)
func DispQueueBFS() {
	fmt.Println("---------------- DispQueueBFS() ------------------")
	tree := dsPkg.Tree{}
	n := 1
	tree.AddNode(n)
	n++

	for i := 0; i < 3; i++ {
		tree.Root.AddNode(n)
		n++
	}

	for i := 0; i < len(tree.Root.Childs); i++ {
		for j := 0; j < 2; j++ {
			tree.Root.Childs[i].AddNode(n)
			n++
		}
	}

	tree.QueuedBFS()
	fmt.Println("\n--------------------------------------------------")
}

func DispPushedHeapTree() {
	fmt.Println("---------------- DispPushedHeapTree() ------------------")
	h := &dsPkg.Heap{}

	h.PushHeapElement(9)
	h.PushHeapElement(8)
	h.PushHeapElement(7)
	h.PushHeapElement(6)
	h.PushHeapElement(5)

	h.PrintHeapElements()
	fmt.Println("--------------------------------------------------------")
}

func DispMinHeapTree() {
	fmt.Println("----------------- DispMinHeapTree() --------------------")
	// [-1, 3, -1, 5, 4]에서 2번째로 큰값 구하기
	h := &dsPkg.Heap{}

	nums := []int{-1, 3, -1, 5, 4}

	for i := 0; i < len(nums); i++ {
		h.PushHeapElement(nums[i])

		if h.GetHeapCount() > 2 {
			h.PopHeapElement()
		}
	}

	fmt.Println(nums)
	fmt.Printf("에서 2번째로 큰값: %d\n", h.PopHeapElement())

	fmt.Println("--------------------------------------------------------")
}

func DispQuizResult(args []int, needle int) {
	fmt.Println("------- DispQuizResult(args []int, needle int) ---------")
	h := &dsPkg.Heap{}

	for i := 0; i < len(args); i++ {
		h.PushHeapElement(args[i])

		if h.GetHeapCount() > needle {
			h.PopHeapElement()
		}
	}

	fmt.Println(args)
	fmt.Printf("에서 %d번째로 큰값: %d\n", needle, h.PopHeapElement())

	fmt.Println("--------------------------------------------------------")
}
