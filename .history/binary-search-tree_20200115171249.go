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

func (t *Tree) insert(data int) {
	if t.root == nil {
		t.root = &Node{value: data}
	} else {
		t.root.insert(data)
	}
}

func (n *Node) insert(data int) {
	if data <= n.value {
		if n.left == nil {
			n.left = &Node{value: data}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{value: data}
		} else {
			n.right.insert(data)
		}
	}
}

func printInOrder(n *Node) {
	if n == nil {
		return
	} else {
		printInOrder(n.left)
		fmt.Printf("%c ", n.value)
		printInOrder(n.right)
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
