package main

import "fmt"

type Node struct {
	left  *Node
	data  interface{}
	right *Node
}

func NewNode(left, right *Node) *Node {
	return &Node{left, nil, right}
}

func (node *Node) SetData(data interface{}) {
	node.data = data
}

func NodeStructs() {
	root := NewNode(nil, nil)
	root.SetData("root node")
	// make child (leaf) nodes:
	left := NewNode(nil, nil)
	left.SetData("left node")
	b := NewNode(nil, nil)
	b.SetData("right node")
	root.left = left
	root.right = b
	fmt.Printf("%v\n", root) //output:&{0xc42000a1a0 root node 0xc42000a1c0}
}
