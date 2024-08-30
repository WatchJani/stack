package stack

import (
	"fmt"
)

const errEmpty = "stack is empty"

type stack[T any] struct {
	store []T
}

func New[T any](capacity int) stack[T] {
	return stack[T]{
		store: make([]T, 0, capacity),
	}
}

func (s *stack[T]) Push(value T) {
	s.store = append(s.store, value)
}

func (s *stack[T]) Clear() {
	s.store = s.store[:0]
}

func (s *stack[T]) Pop() (T, error) {
	if len(s.store) == 0 {
		var zeroValue T
		return zeroValue, fmt.Errorf("%s", errEmpty)
	}
	value := s.store[len(s.store)-1]
	s.store = s.store[:len(s.store)-1]
	return value, nil
}

func (s *stack[T]) IsEmpty() bool {
	return len(s.store) == 0
}

func (s *stack[T]) Peek() (T, error) {
	if len(s.store) == 0 {
		var zeroValue T
		return zeroValue, fmt.Errorf("%s", errEmpty)
	}

	return s.store[len(s.store)-1], nil
}
