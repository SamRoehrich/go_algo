package main

import "fmt"

// Node
type Node struct {
	Key int
	Left *Node
	Right *Node
}

// Insert

func (n *Node) Insert (k int) {
	newNode := &Node{Key: k}
	currentNode := n
	for currentNode != newNode {
		if k > currentNode.Key {
			if currentNode.Right != nil {
				currentNode = currentNode.Right
				} else {
					currentNode.Right = newNode
			}
		} else {
			if currentNode.Left != nil {
				currentNode = currentNode.Left
				} else {
					currentNode.Left = newNode
				}
			}
		}
}

// Search
func (n *Node) Search(k int) bool {
	currentNode := n
	for currentNode != nil {
		if k > currentNode.Key {
			if currentNode.Right != nil {
				currentNode = currentNode.Right
				} else {
					return true
			}
		} else {
			if currentNode.Left != nil {
				currentNode = currentNode.Left
				} else {
					return true
				}
			}
		}
		return false 
}

func InitTree() {
	tree := &Node{Key: 200}
	list := []int {
		100, 200, 300, 11, 123, 12345, 69, 420,
	}
	for _, v := range list {
		tree.Insert(v)
	}
	fmt.Println(tree)
	fmt.Println(*tree.Right)
	fmt.Println(tree.Left)

}