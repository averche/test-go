package main

import "fmt"

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		head: nil,
	}
}

func (l *LinkedList[T]) Add(value T) {
	node := &Node[T]{
		value: value,
		next:  nil,
	}

	if l.head == nil {
		l.head = node
	} else {
		l.head.next = node
	}
}

func (l *LinkedList[T]) ToSlice() []T {
	var slice []T

	n := l.head

	for n != nil {
		slice = append(slice, n.value)
		n = n.next
	}

	return slice
}

func TestLinkedList() {
	list := NewLinkedList[string]()

	list.Add("a")
	list.Add("b")
	list.Add("c")

	fmt.Println(list.ToSlice())
	// Output:
	// [a b c]
}
