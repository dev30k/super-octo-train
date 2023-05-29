package doubleLinkedList

type Node struct {
	Value interface{}
	Next  *Node
	Prev  *Node
}

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) InsertAtBeginning(value interface{}) {
	newNode := &Node{Value: value, Next: ll.Head, Prev: nil}
	ll.Head = newNode

}

func (ll *LinkedList) InsertAtPosition(value interface{}, position int) {

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

func (ll *LinkedList) InsertAtEnd(value interface{}) {
	
}
