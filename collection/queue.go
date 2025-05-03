package collection

import "errors"

type Queue[T any] struct {
	items []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{items: []T{}}
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the front item from the queue. If the queue is empty, returns an error.
func (q *Queue[T]) Dequeue() (T, error) {
	if len(q.items) == 0 {
		var zero T
		return zero, errors.New("queue is empty")
	}
	elem := q.items[0]
	q.items = q.items[1:]
	return elem, nil
}

func (q *Queue[T]) Peek() (T, error) {
	if len(q.items) == 0 {
		var zero T
		return zero, errors.New("queue is empty")
	}
	return q.items[0], nil
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue[T]) Size() int {
	return len(q.items)
}
