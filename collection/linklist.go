package collection

import (
	"fmt"
)

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

type LinkedList[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
}

/* NewLinkedList creates and returns a new, empty linked list. */
func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{Head: nil, Tail: nil}
}

/* Add adds a value to the end of the linked list. */
func (l *LinkedList[T]) Add(value T) {
	node := &Node[T]{Value: value, Next: nil}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
		l.Tail = node
	}
}

/* Remove removes the first occurrence of the given value in the linked list and returns true if it is found, otherwise false. */
func (l *LinkedList[T]) Remove(value T) bool {
	if l.Head == nil {
		return false
	}
	curr := l.Head
	var prev *Node[T] = nil
	for curr != nil {
		if curr.Value == value {
			if curr == l.Head {
				l.Head = curr.Next
			} else if curr == l.Tail {
				prev.Next = nil
				l.Tail = prev
			} else {
				prev.Next = curr.Next
			}
			return true
		}
		prev = curr
		curr = curr.Next
	}
	return false
}

/* Search checks if a given value exists in the linked list and returns true if it does, otherwise false. */
func (l *LinkedList[T]) Search(value T) bool {
	node := l.Head
	for node != nil {
		if node.Value == value {
			return true
		}
		node = node.Next
	}
	return false
}

/* Print prints out the linked list, with each node separated by an arrow (->). */
func (l *LinkedList[T]) Print() {
	node := l.Head
	for node.Next != nil {
		fmt.Print(node.Value, " -> ")
		node = node.Next
	}
	fmt.Println(node.Value)
}
