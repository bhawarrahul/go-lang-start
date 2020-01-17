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

func isExists(node *Node, data int) string {
	if node == nil {
		fmt.Println("Not exists")
	} else if node.value == data {
		printTree(node)
		fmt.Println("Element exists")
	} else {
		isExists(node.left, data)
		isExists(node.right, data)
	}
	return "Not exists"
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
	for _, insertElement := range array {
		// fmt.Println(element)
		tree.makeTree(insertElement)
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
			if tree.root == nil || tree.root.value == 0 {
				fmt.Println("Empty TREE")
			} else {
				printTree(tree.root)
			}
			fmt.Println()
		case 2:
			fmt.Print("Please enter Number of Element: ")
			var elementNumber int = 0
			fmt.Scan(&elementNumber)
			arr := make([]int, elementNumber)
			for count := 0; count < elementNumber; count++ {
				if _, err := fmt.Scan(&arr[count]); err != nil {
					panic(err)
				}
			}
			for _, element := range arr {
				// fmt.Println(element)
				tree.makeTree(element)
			}
		case 3:
			fmt.Print("Please enter Search of Element: ")
			var searchElement int = 0
			fmt.Scan(&searchElement)
			isExists(tree.root, searchElement)
		case 4:
			choice = 4
		}
		if choice == 4 {
			break
		}
	}
}
