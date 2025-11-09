package go_strc

import "sync"

type CircularQueue[T any] struct {
	data        []T
	head, tail  int
	count, size int
	lock        sync.Mutex
}

func NewCircularQueue[T any](size int) *CircularQueue[T] {
	return &CircularQueue[T]{
		data: make([]T, size),
		size: size,
	}
}

func (cq *CircularQueue[T]) Enqueue(value T) bool {
	cq.lock.Lock()
	defer cq.lock.Unlock()

	if cq.count == cq.size {
		return false
	}

	cq.data[cq.tail] = value
	cq.tail = (cq.tail + 1) % cq.size
	cq.count++
	return true
}

func (cq *CircularQueue[T]) Dequeue() (T, bool) {
	cq.lock.Lock()
	defer cq.lock.Unlock()

	if cq.count == 0 {
		var zeroValue T
		return zeroValue, false
	}

	value := cq.data[cq.head]
	cq.head = (cq.head + 1) % cq.size
	cq.count--
	return value, true
}

func (cq *CircularQueue[T]) Peek() (T, bool) {
	cq.lock.Lock()
	defer cq.lock.Unlock()

	if cq.count == 0 {
		var zeroValue T
		return zeroValue, false
	}

	return cq.data[cq.head], true
}

func (cq *CircularQueue[T]) IsEmpty() bool {
	cq.lock.Lock()
	defer cq.lock.Unlock()
	return cq.count == 0
}

func (cq *CircularQueue[T]) IsFull() bool {
	cq.lock.Lock()
	defer cq.lock.Unlock()
	return cq.count == cq.size
}

func (cq *CircularQueue[T]) Size() int {
	cq.lock.Lock()
	defer cq.lock.Unlock()
	return cq.count
}
