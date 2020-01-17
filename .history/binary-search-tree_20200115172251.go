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

func (tree *Tree) insertRootNode(data int) {
	if tree.root == nil {
		tree.root = &Node{value: data}
	} else {
		tree.root.insert(data)
	}
}

func (node *Node) insert(data int) {
	if data <= node.value {
		if node.left == nil {
			node.left = &Node{value: data}
		} else {
			node.left.insert(data)
		}
	} else {
		if node.right == nil {
			node.right = &Node{value: data}
		} else {
			node.right.insert(data)
		}
	}
}

func printBinaryTree(node *Node) {
	if node == nil {
		return
	} else {
		printBinaryTree(node.left)
		fmt.Printf("%c ", node.value)
		printBinaryTree(node.right)
	}
}

func main() {
	var tree Tree

	var array = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for _, element := range array {
		fmt.Println(element)
		tree.insertRootNode(element)
	}

	printBinaryTree(tree.root)
	fmt.Println()
}
