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

func (tree *Tree) insert(data int) {
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

func printBinaryTree(n *Node) {
	if n == nil {
		return
	} else {
		printBinaryTree(n.left)
		fmt.Printf("%c ", n.value)
		printBinaryTree(n.right)
	}
}

func main() {
	var t Tree

	t.insert('F')
	t.insert('B')
	t.insert('A')
	t.insert('D')
	t.insert('C')
	t.insert('E')
	t.insert('G')
	t.insert('I')
	t.insert('H')

	printInOrder(t.root)
	fmt.Println()
}
