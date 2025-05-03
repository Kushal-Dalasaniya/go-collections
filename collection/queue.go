package collection

import "errors"

type Queue[T any] struct {
	items []T
}

/* NewQueue creates and returns a new instance of Queue with an empty slice of items. */
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{items: []T{}}
}

/* Enqueue adds an item to the end of the queue. */
func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

/* Dequeue removes and returns the front item from the queue. If the queue is empty, returns an error. */
func (q *Queue[T]) Dequeue() (T, error) {
	if len(q.items) == 0 {
		var zero T
		return zero, errors.New("queue is empty")
	}
	elem := q.items[0]
	q.items = q.items[1:]
	return elem, nil
}

/* Peek returns the front item from the queue without removing it. If the queue is empty, returns an error. */
func (q *Queue[T]) Peek() (T, error) {
	if len(q.items) == 0 {
		var zero T
		return zero, errors.New("queue is empty")
	}
	return q.items[0], nil
}

/* IsEmpty checks whether the queue is empty and returns true if it is, otherwise false. */
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

/* Size returns the number of items in the queue. */
func (q *Queue[T]) Size() int {
	return len(q.items)
}
