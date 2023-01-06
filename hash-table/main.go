package main

import "fmt"

/*
Implementing a hash table usign linked list
to resolve collisions.
*/

// The size of the hashtable array
const ArraySize = 7

// HashTable structure
// HashTable holds an array
type HashTable struct {
	array [ArraySize]*bucket
}

// Bucket structure
type bucket struct {
	head *bucketNode
}

// Bucket Node structure
// bucketNode is a linked list
type bucketNode struct {
	key  string
	next *bucketNode
}

// HashTable
// Insert will take in a key and add it to the hash table array
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].Insert(key)
}

// Search will take a key and return true if the key is stored in the hash table
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].Search(key)
}

// Delete will take in a key and delete it from the hash table
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].Delete(key)
}

// Bucket

// Insert will take a key, create a node with the key and insert the node in the bucket
func (b *bucket) Insert(k string) {
	if !b.Search(k) {
		newNode := &bucketNode{key: k}
		// Link the current head to the new node
		newNode.next = b.head
		// Make the new node the head
		b.head = newNode
	} else {
		fmt.Printf("%s already exists \n", k)
	}
}

// Search will take in a key, loop through the bucket and return true if a match was found
func (b *bucket) Search(k string) bool {
	currentNode := b.head
	if currentNode == nil {
		return false
	}
	if currentNode.key == k {
		return true
	}
	for currentNode.next != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// Delete will take in a key and delete the node from the bucket
func (b *bucket) Delete(k string) {
	currentNode := b.head

	if currentNode.key == k {
		b.head = currentNode.next
		return
	}

	for currentNode.next != nil {
		if currentNode.next.key == k {
			currentNode.next = currentNode.next.next
		}
		currentNode = currentNode.next
	}
}

// Hash function
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	// return the remainder of sum/arraySize
	return sum % ArraySize
}

// Init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	hashTable := Init()
	list := []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOLKEN",
	}
	for _, v := range list {
		hashTable.Insert(v)
	}
	fmt.Println(hashTable.Search("STAN"))
	hashTable.Delete("STAN")
	fmt.Println(hashTable.Search("STAN"))
	fmt.Println(hashTable.Search("TOLKEN"))
	fmt.Println(hashTable.Search("PETER"))
}
