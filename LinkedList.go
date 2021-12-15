package main

type Node struct {
	value int
	next  *Node
}

type List struct {
	head *Node
	tail *Node
}

func createLinkedList() List {
	list := List{}
	head := Node{}
	list.head = &head
	list.tail = &head
	return list
}

func insert(list *List, number int) {
	node := Node{
		value: number,
	}
	list.tail.next = &node
	list.tail = &node
}
