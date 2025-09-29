package main

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() T {
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]

	return element
}

func (s *Stack[T]) Peek() T {
	return s.elements[len(s.elements)-1]
}

func (s *Stack[T]) Size() int {
	return len(s.elements)
}

func (s *Stack[T]) Any() bool {
	return len(s.elements) > 0
}
