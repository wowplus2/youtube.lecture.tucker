// BST: Binary Search Tree
package main

import (
	"fmt"

	dsPkg "lec.youtube.tucker/bst/dataStruct"
)

func DispBinaryTree(t *dsPkg.BinaryTree) {
	fmt.Println("\n--------- DispBinaryTree(t *dsPkg.BinaryTree) ----------")
	t.Root.AddNode(3)
	t.Root.AddNode(2)
	t.Root.AddNode(4)
	t.Root.AddNode(8)
	t.Root.AddNode(7)
	t.Root.AddNode(6)
	t.Root.AddNode(10)
	t.Root.AddNode(9)

	t.DispBinaryTree()
	fmt.Println("\n--------------------------------------------------------")
}

func DispSearchAtBinaryTree(t *dsPkg.BinaryTree, n int) {
	fmt.Println("-- DispSearchAtBinaryTree(t *dsPkg.BinaryTree, n int) --")
	if found, cnt := t.SearchBinaryTree(n); found {
		fmt.Printf("found %d / cnt:%d\n", n, cnt)
	} else {
		fmt.Printf("not found %d / cnt: %d\n", n, cnt)
	}
	fmt.Println("--------------------------------------------------------")
}

func main() {
	iniNum := 5 // tree 최상단 값(기준값)
	fndNum := 6 // 찾을 값

	tree := dsPkg.NewBinaryTree(iniNum)

	DispBinaryTree(tree)
	DispSearchAtBinaryTree(tree, fndNum)
}
