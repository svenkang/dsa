package main

import "fmt"

type Node struct {
	Value int
	Left *Node
	Right *Node
}

func (n *Node) Insert (v int) {
	if n.Value < v {
		// move right
		if n.Right == nil {
			n.Right = &Node{Value: v}
		} else {
			n.Right.Insert(v)
		}
	} else if n.Value > v {
		// move left
		if n.Left == nil {
			n.Left = &Node{Value: v}
		} else {
			n.Left.Insert(v)
		}
	}
}

func (n *Node) Search (v int) bool {
	if n == nil {
		return false
	}

	if n.Value < v {
		return n.Right.Search(v)
	} else if n.Value > v {
		return n.Left.Search(v)
	}

	return true
}

func main() {
	tree := &Node{Value: 3}
	tree.Insert(5)
	tree.Insert(8)
	fmt.Println(tree.Search(23))
}
