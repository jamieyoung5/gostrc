package go_strc

import (
	"container/list"
	"sync"
)

type Stack[T any] struct {
	list *list.List
	lock sync.Mutex
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{list: list.New()}
}

func (s *Stack[T]) Push(value T) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.list.PushBack(value)
}

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

func (s *Stack[T]) IsEmpty() bool {
	return s.list.Len() == 0
}

func (s *Stack[T]) Size() int {
	return s.list.Len()
}
