package collection

import (
	"errors"
	"fmt"
)

type Stack[T any] struct {
	items []T
}

/* NewStack creates and returns a new instance of Stack with an empty slice of items. */
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{items: []T{}}
}

/* Push adds an item to the top of the stack. */
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

/* Pop removes the top item from the stack and returns it. If the stack is empty, returns an error. */
func (s *Stack[T]) Pop() (T, error) {
	if len(s.items) == 0 {
		var zero T
		return zero, errors.New("stack is empty")
	}
	elem := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return elem, nil
}

/* Peek returns the top item from the stack without removing it. If the stack is empty, returns an error. */
func (s *Stack[T]) Peek() (T, error) {
	if len(s.items) == 0 {
		var zero T
		return zero, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

/* IsEmpty checks whether the stack is empty and returns true if it is, otherwise false. */
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

/* Size returns the number of items in the stack. */
func (s *Stack[T]) Size() int {
	return len(s.items)
}

/* Print prints out the stack, with each item separated by a space. */
func (s *Stack[T]) Print() {
	fmt.Printf("%v\n", s.items)
}
