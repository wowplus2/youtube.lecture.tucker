package main

import "fmt"

type Node struct {
	next *Node
	val  int
}

func dispNodes(root *Node) {
	node := root
	for node.next != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.next
	}
	fmt.Printf("%d\n", node.val)
}

func addNodeFocusAppend(root *Node, val int) {
	var tail *Node
	tail = root
	for tail.next != nil {
		tail = tail.next
	}

	//새로운 node 추가
	node := &Node{val: val}
	tail.next = node
}

func linkedListAdd01() {
	var root *Node

	root = &Node{val: 0}

	for i := 1; i < 10; i++ {
		addNodeFocusAppend(root, i)
	}

	dispNodes(root)
}

func addNodeFocusTail(tail *Node, val int) *Node {
	//새로운 node 추가
	node := &Node{val: val}
	tail.next = node

	return node
}

func linkedListAdd02() {
	var root *Node
	var tail *Node

	root = &Node{val: 0}
	tail = root

	for i := 1; i < 10; i++ {
		tail = addNodeFocusTail(tail, i)
	}

	dispNodes(root)
}

func removeNodeMid(prev *Node) {
	prev.next = prev.next.next
}

func removeNodeMulti(node *Node, root *Node, tail *Node) (*Node, *Node) {
	if node == root {
		root = root.next

		if root == nil {
			tail = nil
		}

		return root, tail
	}

	prev := root
	for prev.next != node {
		prev = prev.next
	}

	if node == tail {
		prev.next = nil
		tail = prev
	} else {
		prev.next = prev.next.next
	}

	return root, tail
}

func linkedListRemove01() {
	var root *Node
	var tail *Node

	root = &Node{val: 0}
	tail = root

	for i := 1; i < 10; i++ {
		tail = addNodeFocusTail(tail, i)
	}

	dispNodes(root)
	removeNodeMid(root) // linkedList 안에 값을 삭제
	dispNodes(root)
}

func linkedListRemove02() {
	var root *Node
	var tail *Node

	root = &Node{val: 0}
	tail = root

	for i := 1; i < 10; i++ {
		tail = addNodeFocusTail(tail, i)
	}

	dispNodes(root)
	root, tail = removeNodeMulti(root, root, tail) // linkedList의 root 값을 삭제
	dispNodes(root)
	root, tail = removeNodeMulti(tail, root, tail) // linkedList의 tail 값을 삭제
	dispNodes(root)
}

func main() {
	//linkedListAdd01()	// 맨뒤에 node를 계속 추가하는 방법(n개의 node를 추가하 기위해서 for가 n번 순회한다.) - O(N)
	//linkedListAdd02()	// tail을 기록하고 있으므로 node에 tail을 추가하는 방법(n개의 node를 추가하 기위해서 for가 1번 순회한다.) - O(1)
	//linkedListRemove01()
	linkedListRemove02() // remove하고자 하는 node까지 n번 순회후 해당 node를 삭제한다. - O(N)
}
