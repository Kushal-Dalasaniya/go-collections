package collection

type empty struct{}

type Set[T comparable] struct {
	items map[T]empty
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{items: make(map[T]empty)}
}

func (s *Set[T]) Add(item T) {
	s.items[item] = empty{}
}

func (s *Set[T]) Remove(item T) {
	delete(s.items, item)
}

func (s *Set[T]) Contains(item T) bool {
	_, exists := s.items[item]
	return exists
}

func (s *Set[T]) Size() int {
	return len(s.items)
}

func (s *Set[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Set[T]) Clear() {
	s.items = make(map[T]empty)
}

func (s *Set[T]) Items() []T {
	res := make([]T, 0, len(s.items))
	for k := range s.items {
		res = append(res, k)
	}
	return res
}
