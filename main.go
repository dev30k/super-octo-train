package main

import (
	"docker-test/singleLinkedList"
	"fmt"
)

func main() {
	// Create a new instance of LinkedList
	ll := &singleLinkedList.LinkedList{}

	// Insert nodes
	ll.InsertAtBeginning("Hello")
	ll.InsertAtBeginning("!")

	// Traverse and print the linked list
	ll.TraverseLinkedList(func(value interface{}) {
		fmt.Print(value)
	})

}
