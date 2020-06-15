package main

import "fmt"

type Node struct {
	next *Node
	prev *Node
	val  int
}

type LinkedList struct {
	root *Node
	tail *Node
}

func (l *LinkedList) addNode(n int) {
	if l.root == nil {
		l.root = &Node{val: n}
		l.tail = l.root
		return
	}

	l.tail.next = &Node{val: n}
	prev := l.tail
	l.tail = l.tail.next
	l.tail.prev = prev
}

func (l *LinkedList) removeNode(node *Node) {
	if node == l.root {
		l.root = l.root.next
		l.root.prev = nil
		node.next = nil
		return
	}

	prev := node.prev
	if node == l.tail {
		prev.next = nil
		l.tail.prev = nil
		l.tail = prev
	} else {
		node.prev = nil
		prev.next = prev.next.next
		prev.next.prev = prev
	}

	node.next = nil
}

func (l *LinkedList) printNodes() {
	node := l.root
	for node.next != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.next
	}
	fmt.Printf("%d\n", node.val)
}

func (l *LinkedList) printReverseNodes() {
	node := l.tail
	for node.prev != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.prev
	}

	fmt.Printf("%d\n", node.val)
}

func main() {
	list := &LinkedList{}
	list.addNode(0) // root node를 추가한다.

	for i := 1; i < 10; i++ {
		list.addNode(i) // linked-list에 node를 추가한다.
	}

	list.printNodes()
	list.removeNode(list.root.next) // root 다음 node를 삭제한다.
	list.printNodes()
	list.removeNode(list.root) // root node를 삭제한다.
	list.printNodes()
	list.removeNode(list.tail) // tail node를 삭제한다.
	list.printNodes()

	fmt.Printf("tail: %d\n", list.tail.val)

	list.printReverseNodes()

	a := []int{1, 2, 3, 4, 5}
	a = append(a[0:2], a[3:]...)
	fmt.Println(a)
}
