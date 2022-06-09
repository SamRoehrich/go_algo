package main

import (
	"fmt"
)

const ArraySize = 7

// Hash table = array of linked lists
type HashTable struct {
	array [ArraySize]*bucket
}

// Insert - create and insert new node into hash map
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Search - Compute the hash and traverse the list at index
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// Delete - Compute hash, delete node and re-link list
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// ----------------------------------------

// Bucket - Start of the linked list
type bucket struct {
	head *bucketNode
}

// bucketNode - each node of the list
type bucketNode struct {
	key string
	next *bucketNode
}

func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k , "already exists in the data structure.")
	}
}

func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (b *bucket) delete(k string) {
	// if k is the current head, move the head down one
	if b.head.key == k {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	// current node and not the end of the list
	for previousNode.next != nil {
		// next node.key == key
		if previousNode.next.key == k {
			// set current node.next to be deleted.next
			previousNode.next = previousNode.next.next
		}
	}
}


// hashing function
func hash(key string) int {
	sum := 0
	// loop over the values of key
	for _,v := range key{
		// add the int value of key to sum
		sum += int(v)
	}
	
	return sum % ArraySize
}

// Init - add a bucket to each index of the Hashmap.array
func Init() *HashTable {
	result := HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return &result
}

func hashMap () {
	testHashTable := Init()
	list1 := []string{
		"Sam",
		"Jandey",
		"John",
		"Anthony",
		"Brad",
		"Jon",
		"Jenna",
	}

	for _,v := range list1 {
		testHashTable.Insert(v)
	}

	testHashTable2 := Init()
	list2 := []string {
		"Sam",
		"Sam",
		"John",
		"John",
		"Luke",
		"Jason",
		"Myra",
	}
	for _,v := range list2 {
		testHashTable2.Insert(v)
	}
}