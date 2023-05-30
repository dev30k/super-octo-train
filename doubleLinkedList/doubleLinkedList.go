package doubleLinkedList

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
	Prev  *Node
}

type DoubleLinkedList struct {
	Head *Node
	Tail *Node
}

func (ll *DoubleLinkedList) InsertAtBeginning(value interface{}) {
	newNode := &Node{Value: value, Prev: nil, Next: ll.Head}
	if ll.Head != nil {
		ll.Head.Prev = newNode
	}
	ll.Head = newNode
	if ll.Tail == nil {
		ll.Tail = newNode
	}
}
func (ll *DoubleLinkedList) TraverseLinkedList(fn func(value interface{})) {
	currentNode := ll.Head
	for currentNode.Next != nil {
		fn(currentNode.Value)
		fmt.Println()
		currentNode = currentNode.Next

	}
}
func (ll *DoubleLinkedList) InsertAtPosition(value interface{}, position int) {

	newNode := &Node{Value: value}
	current := ll.Head
	count := 0

	if position == 0 {
		if ll.Head != nil {
			ll.Head.Prev = newNode

		}
		newNode.Next = ll.Head
		ll.Head = newNode
		return

	}
	if current == nil {
		return
	}

	for current != nil && count < position-1 {
		current = current.Next
		count++
	}

	newNode.Next = current.Next
	newNode.Prev = current

	if current.Next != nil {
		current.Next.Prev = newNode
	}

	current.Next = newNode

}
func (ll *DoubleLinkedList) InsertAtEnd(value interface{}) {
	newNode := &Node{Value: value}
	current := ll.Head

	for current.



}
