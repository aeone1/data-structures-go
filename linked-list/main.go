package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	// make a temp space called second
	second := l.head
	// set the new node as the head
	l.head = n
	// let the new head point to the old head
	l.head.next = second
	// increase the length
	l.length++
}

// To print out the values of a linked list
// we can use a value reciever
func (l linkedList) printListData() {
	// store the head to print
	toPrint := l.head
	// print until there is no more node left
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		// update to next node in the list
		toPrint = toPrint.next
		// decrease length by 1
		l.length--
	}
	fmt.Println()
}

func (l *linkedList) deleteWithValue(value int) {
	// Return early for empty list
	if l.length == 0 {
		return
	}

	// If the value to delete is the header
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	// Since this is a single-linked-list
	// we can only see the next node and cannot
	// see the previous node(we need this to link it
	// to the next.next node, ie skip/delete a node)
	// once we've found the value to delete, we need to work
	// with the current node's next node
	currentNode := l.head
	// Loop till we find a next node who's value
	// matches what we want to delete
	for currentNode.next.data != value {
		// Return(not found) if we are at the end of the linkedlist
		if currentNode.next.next == nil {
			return
		}
		// Go to next node
		currentNode = currentNode.next
	}
	// we've probably found the node to delete by its value
	// So we delete it by skipping it
	currentNode.next = currentNode.next.next
	// reduce the length of l
	l.length--

}

func main() {
	myList := linkedList{}
	node1 := &node{data: 10}
	node2 := &node{data: 20}
	node3 := &node{data: 30}
	node4 := &node{data: 20}

	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)

	myList.deleteWithValue(100)
	myList.printListData()
	myList.deleteWithValue(20)
	myList.printListData()
	myList.deleteWithValue(10)
	myList.printListData()

	emptyList := linkedList{}
	emptyList.deleteWithValue(12)
	emptyList.printListData()

}
