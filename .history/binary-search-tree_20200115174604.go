package main

import "fmt"

type Tree struct {
	root *Node
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func (tree *Tree) makeTree(data int) {
	if tree.root == nil {
		tree.root = &Node{value: data}
	} else {
		tree.root.createChildNode(data)
	}
}

func (node *Node) createChildNode(data int) {
	if data <= node.value {
		if node.left == nil {
			// fmt.Println("in left child Node")
			node.left = &Node{value: data}
		} else {
			node.left.createChildNode(data)
		}
	} else {
		// fmt.Println("in Right child Node")
		if node.right == nil {
			node.right = &Node{value: data}
		} else {
			node.right.createChildNode(data)
		}
	}
	// fmt.Println(data)
}

func printBinaryTree(node *Node) {
	if node == nil {
		return
	} else {
		printBinaryTree(node.left)
		fmt.Print(" Node:-> ")
		fmt.Print(node.value)
		printBinaryTree(node.right)
	}
}

func main() {
	var tree Tree

	var array = []int{19, 1, 2, 16, 32, 64, 12, 4, 8}
	for _, element := range array {
		// fmt.Println(element)
		tree.makeTree(element)
	}

	printTree(tree.root)
	fmt.Println()
}
