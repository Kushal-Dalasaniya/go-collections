package collection

type empty struct{}

type Set[T comparable] struct {
	items map[T]empty
}

/*
NewSet creates and returns a new instance of Set, initialized with an empty map.
The Set is designed to hold unique elements of any type that implements the comparable interface.
*/
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{items: make(map[T]empty)}
}

/* Add inserts the specified item into the set. If the item already exists in the set, it is ignored. */
func (s *Set[T]) Add(item T) {
	s.items[item] = empty{}
}

/* Remove deletes the specified item from the set. If the item does not exist in the set, it is ignored. */
func (s *Set[T]) Remove(item T) {
	delete(s.items, item)
}

/* Contains checks if a given item exists in the set and returns true if it does, otherwise false. */
func (s *Set[T]) Contains(item T) bool {
	_, exists := s.items[item]
	return exists
}

/* Size returns the number of items in the set. */
func (s *Set[T]) Size() int {
	return len(s.items)
}

/* IsEmpty checks if the set is empty and returns true if it is, otherwise false. */
func (s *Set[T]) IsEmpty() bool {
	return len(s.items) == 0
}

/* Clear removes all items from the set, returning the set to its initial state. */
func (s *Set[T]) Clear() {
	s.items = make(map[T]empty)
}

/* Items returns a slice of all items in the set, in the order they were added. */
func (s *Set[T]) Items() []T {
	res := make([]T, 0, len(s.items))
	for k := range s.items {
		res = append(res, k)
	}
	return res
}
