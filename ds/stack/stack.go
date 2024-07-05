package ds

import (
	"errors"
)

var (
	EmptyStackErr = errors.New("Cannot operate on an empty stack")
)

type (
	node[T any] struct {
		value T
		prev *node[T]
	}
	Stack[T any] struct {
		top *node[T]
		size uint
	}
)

func (s *Stack[T]) Push(v T) T {
	s.top = &node[T]{ value: v, prev: s.top }
	s.size++
	return v
}

func (s *Stack[T]) Pop() (T, error) {
	if s.size == 0 {
		return *new(T), EmptyStackErr
	}
	v := s.top.value 
	s.top = s.top.prev
	s.size--
	return v, nil
}

func (s *Stack[T]) Peek() (T, error) {
	if s.size == 0 {
		return *new(T), EmptyStackErr
	}
	return s.top.value, nil
}

func (s *Stack[T]) Len() uint {
	return s.size
}
