package main

import (
	"fmt"
)

type binaryTree struct {
	value int
	pre   *binaryTree
	next  *binaryTree
}

func addPre(tree *binaryTree, value int) (binaryTree1 binaryTree) {
	binaryTree1.value = value
	tree.pre = &binaryTree1
	return binaryTree1
}

func addNext(tree *binaryTree, value int) binaryTree {
	binaryTree1 := new(binaryTree)
	binaryTree1.value = value
	tree.next = binaryTree1
	return *binaryTree1
}

func structsBinaryTree() {
	binaryTree1 := new(binaryTree)
	binaryTree1.value = 10
	pre := addPre(binaryTree1, 9)
	next := addNext(binaryTree1, 11)
	fmt.Printf("The node is %d. The pre is %d. The next is %d\n",
		binaryTree1.value, pre.value, next.value)
}
