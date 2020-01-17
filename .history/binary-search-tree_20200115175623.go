package main

import (
	"bufio"
	"fmt"
	"os"
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
	reader := bufio.NewReader(os.Stdin)
	for {
		switch choice {
		case 1:
			printTree(tree.root)
			fmt.Println()
		case 2:
			fmt.Println("Bonjour")
		case 3:
			fmt.Println("Namstay")
		default:

		}
	}

}
