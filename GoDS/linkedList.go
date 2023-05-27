//linked list go 

package gods


type Node struct{
	Value interface{}
	Next *Node
}

type LinkedList struct{
	Head *Node
}

func (ll *LinkedList) InsertAtBeginning(value interface{}){
	newNode  := &Node{Value: value,Next: ll.Head}
	ll.Head = newNode

}
func (ll *LinkedList) InsertAtEnd(value interface{}){
	newNode := &Node{Value: value}

	if ll.Head == nil{
		ll.Head = newNode
		return
	}
	current := ll.Head
	for current.Next != nil{
		current = current.Next
	}
	current.Next = newNode


}

func (ll *LinkedList) InsertAtPosition(value interface{}, position int){

	newNode := &Node{Value: value}

	if position <= 0{
		newNode.Next = ll.Head
		ll.Head = newNode
		return
	}

	current := ll.Head
	prev := ll.Head
	count := 0
	for current != nil && count < position{
		prev = current
		current = current.Next
		count++
	}
	newNode.Next = current
    prev.Next = newNode
}

func (ll *LinkedList) TraverseLinkedList(){
	current := ll.Head

	if current != nil{
		for current != nil{
			current = current.Next
		}
	}else{
		panic("List is empty")
	}

}

func (ll *LinkedList) DeleteBeginning(value interface{}){
	if ll.Head == nil{
		return 
	}
	ll.Head = ll.Head.Next

}

func (ll *LinkedList) DeleteEnd(value interface{}){
	if ll.Head.Next == nil{
		return 
		
	}
	current := ll.Head

	for current != nil{
		
	}


}