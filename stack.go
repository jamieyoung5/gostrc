package gostrc

import (
	"container/list"
	"sync"
)

// Stack is an ordered set of elements where only the first element of the set can be changed or viewed.
// You can: push (add a new element to the top of the stack i.e becomes the first element of the set),
// pop (removes the element currently at the top of the stack) and peek (get the top element of the stack)
// https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
type Stack[T any] struct {
	list *list.List
	lock sync.Mutex
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{list: list.New()}
}

// Push adds a new element to the top of the stack
func (s *Stack[T]) Push(value T) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.list.PushBack(value)
}

// Pop removes the element at the top of the stack
func (s *Stack[T]) Pop() T {
	s.lock.Lock()
	defer s.lock.Unlock()

	element := s.list.Back()
	if element != nil {
		s.list.Remove(element)
		return element.Value.(T)
	}

	var zeroValue T
	return zeroValue
}

// Peek returns the element at the top of the stack
func (s *Stack[T]) Peek() T {
	s.lock.Lock()
	defer s.lock.Unlock()

	element := s.list.Back()
	if element != nil {
		return element.Value.(T)
	}

	var zeroValue T
	return zeroValue
}

// IsEmpty returns true if the stack has no elements, false if it does
func (s *Stack[T]) IsEmpty() bool {
	return s.list.Len() == 0
}

// Size returns the number of elements in the stack
func (s *Stack[T]) Size() int {
	return s.list.Len()
}
