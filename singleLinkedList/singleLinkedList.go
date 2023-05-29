//single linked list go implementation

package singleLinkedList

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) InsertAtBeginning(value interface{}) {
	newNode := &Node{Value: value, Next: ll.Head}
	ll.Head = newNode

}
func (ll *LinkedList) InsertAtEnd(value interface{}) {
	newNode := &Node{Value: value}

	if ll.Head == nil {
		ll.Head = newNode
		return
	}
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode

}

func (ll *LinkedList) InsertAtPosition(value interface{}, position int) {

	newNode := &Node{Value: value}

	if position <= 0 {
		newNode.Next = ll.Head
		ll.Head = newNode
		return
	}

	current := ll.Head
	prev := ll.Head
	count := 0
	for current != nil && count < position {
		prev = current
		current = current.Next
		count++
	}
	newNode.Next = current
	prev.Next = newNode
}

func (ll *LinkedList) TraverseLinkedList(fn func(value interface{})) {
	current := ll.Head
	for current != nil {
		fn(current.Value)
		fmt.Println() // Add a newline character after printing each value
		current = current.Next
	}
}

func (ll *LinkedList) DeleteBeginning() {
	if ll.Head == nil {
		return
	}
	ll.Head = ll.Head.Next

}

func (ll *LinkedList) DeleteEnd() {
	if ll.Head.Next == nil {
		return

	}
	if ll.Head.Next == nil {
		ll.Head = nil
		return
	}
	current := ll.Head

	for current.Next.Next != nil {
		current = current.Next

	}
	current.Next = nil

}

func (ll *LinkedList) DeletePosition(position int) {
	if ll.Head.Next == nil {
		return
	}
	if position < 0 {
		return
	}
	if position == 0 {
		ll.Head = ll.Head.Next
		return
	}

	current := ll.Head
	prev := ll.Head
	count := 0

	for current != nil && count < position {
		prev = current
		current = current.Next
		count++
	}
	if current == nil {
		return
	}
	prev.Next = current.Next

}
