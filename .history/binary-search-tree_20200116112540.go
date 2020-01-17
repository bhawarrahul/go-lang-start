package main

import (
	"fmt"
)

type BSTTree struct {
	root *Node
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func (tree *BSTTree) makeTree(data int) {
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

func isExists(node *Node, element int) {
	if node == nil {
		fmt.Println("Not exists")
	} else if node.value == element {
		fmt.Println("Element exists")
		fmt.Print("Sub Tress: ")
		printTree(node)
		fmt.Println()
	} else {
		if node.left != nil {
			isExists(node.left, element)
		}
		if node.right != nil {
			isExists(node.right, element)
		}
	}
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
	var tree BSTTree

	var array = []int{19, 1, 2, 16, 32, 64, 12, 4, 8}
	for _, insertElement := range array {
		// fmt.Println(element)
		tree.makeTree(insertElement)
	}

	var choice int = 4

	for {
		fmt.Println("1. Print Tree")
		fmt.Println("2. Insert Tree Element")
		fmt.Println("3. Is Exist")
		fmt.Println("4. Exit")
		fmt.Print("Enter Choice: ")

		fmt.Scanln(&choice)
		if choice == 4 {
			break
		}
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
		default:
			fmt.Println("Invalid Choice")
		}
	}
}
