package main

import "fmt"

type Node struct {
	value int
	point *Node
}

func addNextNode(node *Node, value int) {
	for true {
		if newNode := node.point; newNode == nil {
			newNode := new(Node)
			newNode.value = value
			node.point = newNode
			return
		}
		node = node.point
	}
}

func structsNode() {
	node := new(Node)
	node.value = 10
	addNextNode(node, 9)
	addNextNode(node, 8)
	addNextNode(node, 7)
	addNextNode(node, 6)
	fmt.Printf("The start node value is %d!\n", node.value)
	fmt.Printf("The second node value is %d!\n", node.point.value)
	fmt.Println("========start print node=========")
	for true {
		fmt.Printf("The node value is %d!\n", node.value)
		if newNode := node.point; newNode != nil {
			fmt.Printf("The node value is %d!\n", newNode.value)
			node = node.point
		} else {
			break
		}
	}
	fmt.Println("Over")
}
