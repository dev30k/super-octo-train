package doubleLinkedList_test

import (
	"docker-test/doubleLinkedList"
	"fmt"
	"testing"
)

// there is a problem here that i cannot see
func TestLinkedList_InsertAtBeginning(t *testing.T) {
	ll := &doubleLinkedList.DoubleLinkedList{}

	ll.InsertAtBeginning("World")
	ll.InsertAtBeginning("Hello")
	ll.InsertAtBeginning("Hello")
	values := make([]string, 0)
	ll.TraverseLinkedList(func(value interface{}) {

		values = append(values, value.(string))
		fmt.Println(values)
	})
	fmt.Println(values)
	expected := []string{"Hello", "Hello"}
	if len(values) != len(expected) {
		t.Errorf("Unexpected linked list length: got %d, want %d", len(values), len(expected))
	}

	for i := range values {
		if values[i] != expected[i] {
			t.Errorf("Unexpected value at position %d: got %s, want %s", i, values[i], expected[i])
		}
	}

}
