package main

import "testing"

func TestShouldReturnLinkedList(t *testing.T) {
	list := createLinkedList()
	for i := 1; i < 10; i++ {
		insert(&list, i)
	}
	current := list.head
	counter := 0
	for current.next != nil {
		if current.value != counter {
			t.Error("Fail", counter, current.value)
		}
		counter++
		current = current.next
	}

}
