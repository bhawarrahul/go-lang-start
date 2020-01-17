package main

import (
	"fmt"
)

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

func printTree(node *Node) {
	if node == nil {
		return
	} else {
		printTree(node.left)
		fmt.Print(" Node:-> ")
		fmt.Print(node.value)
		printTree(node.right)
	}
}

func main() {
	var tree Tree

	var array = []int{19, 1, 2, 16, 32, 64, 12, 4, 8}
	for _, element := range array {
		// fmt.Println(element)
		tree.makeTree(element)
	}

	var choice int = 1

	for {
		fmt.Println("1. Print Tree")
		fmt.Println("2. Insert Tree Element")
		fmt.Println("3. Is Exist")
		fmt.Println("4. Exit")

		fmt.Scanln(&choice)
		switch choice {
		case 1:
			printTree(tree.root)
			fmt.Println()
		case 2:
			fmt.Println("Please enter Number of Element")
			var elementNumber int = 0
			Scan(&elementNumber)
			for count := 0; count < elementNumber; count++ {

			}
		case 3:
			fmt.Println("Is Exist")
		case 4:
			choice = 4
		}
		if choice == 4 {
			break
		}
	}

}
