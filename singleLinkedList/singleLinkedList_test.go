package singleLinkedList_test

import (
	"docker-test/singleLinkedList"
	"testing"
)

func TestLinkedList_InsertAtBeginning(t *testing.T) {
	ll := &singleLinkedList.LinkedList{}

	ll.InsertAtBeginning("World")
	ll.InsertAtBeginning("Hello")

	values := make([]string, 0)
	ll.TraverseLinkedList(func(value interface{}) {
		values = append(values, value.(string))
	})
	expected := []string{"Hello", "World"}
	if len(values) != len(expected) {
		t.Errorf("Unexpected linked list length: got %d, want %d", len(values), len(expected))
	}

	for i := range values {
		if values[i] != expected[i] {
			t.Errorf("Unexpected value at position %d: got %s, want %s", i, values[i], expected[i])
		}
	}
}

func TestLinkedList_InsertAtPosition(t *testing.T) {
	ll := &singleLinkedList.LinkedList{}

	ll.InsertAtBeginning("e")
	ll.InsertAtBeginning("m")
	ll.InsertAtBeginning("e")
	ll.InsertAtBeginning("G")

	//test for insert at position
	ll.InsertAtPosition("n", 2)
	ll.InsertAtPosition("o", 3)

	values := make([]string, 0)
	ll.TraverseLinkedList(func(value interface{}) {
		values = append(values, value.(string))
	})

	expected := []string{"G", "e", "n", "o", "m", "e"}
	if len(values) != len(expected) {
		t.Errorf("Unexpected linked list length: got %d, want %d", len(values), len(expected))
	}
	for i := range values {
		if values[i] != expected[i] {
			t.Errorf("Unexpected value at position %d: got %s, want %s", i, values[i], expected[i])
		}
	}
}
