package stack

import (
	"fmt"
)

const errEmpty = "stack is empty"

type Stack[T any] struct {
	store []T
}

func New[T any](capacity int) Stack[T] {
	return Stack[T]{
		store: make([]T, 0, capacity),
	}
}

func (s *Stack[T]) Push(value T) {
	s.store = append(s.store, value)
}

func (s *Stack[T]) Clear() {
	s.store = s.store[:0]
}

func (s *Stack[T]) Pop() (T, error) {
	if len(s.store) == 0 {
		var zeroValue T
		return zeroValue, fmt.Errorf("%s", errEmpty)
	}
	value := s.store[len(s.store)-1]
	s.store = s.store[:len(s.store)-1]
	return value, nil
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.store) == 0
}

func (s *Stack[T]) Peek() (T, error) {
	if len(s.store) == 0 {
		var zeroValue T
		return zeroValue, fmt.Errorf("%s", errEmpty)
	}

	return s.store[len(s.store)-1], nil
}
