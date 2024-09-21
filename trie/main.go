package main

import "fmt"

const AlphabetSize = 26

// Node
type Node struct {
	children [AlphabetSize]*Node
	isEnd bool
}

// Trie
type Trie struct {
	root *Node
}

// Init
func InitTrie() *Trie {
	return &Trie{root: &Node{}}
}

// Insert 
func (t *Trie) Insert(w string) {
	currentNode := t.root
	for i := 0; i < len(w); i ++ {
		charIndex := w[i] - 'a' 
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

// Search
func (t *Trie) Search(w string) bool {
	currentNode := t.root
	for i := 0; i < len(w); i ++ {
		charIndex := w[i] - 'a' 
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd {
		return true
	}
	return false
}

func main() {
	t := InitTrie()
	t.Insert("aragorn")
	t.Insert("aragog")
	fmt.Println(t.Search("aragog"))
}
