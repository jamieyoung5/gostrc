package gostrc

import "sync"

// CircularQueue implements a thread-safe fixed-size fifo ring buffer
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

// Enqueue adds a new element to the back of the queue
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

// Dequeue removes the element at the front of the queue
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

// Peek returns the element at the head of the queue, if there is one
func (cq *CircularQueue[T]) Peek() (T, bool) {
	cq.lock.Lock()
	defer cq.lock.Unlock()

	if cq.count == 0 {
		var zeroValue T
		return zeroValue, false
	}

	return cq.data[cq.head], true
}

// IsEmpty returns true if the queue is empty, false if not
func (cq *CircularQueue[T]) IsEmpty() bool {
	cq.lock.Lock()
	defer cq.lock.Unlock()
	return cq.count == 0
}

// IsFull returns true if the queue is full, false if not
func (cq *CircularQueue[T]) IsFull() bool {
	cq.lock.Lock()
	defer cq.lock.Unlock()
	return cq.count == cq.size
}

// Size returns the number of queued elements
func (cq *CircularQueue[T]) Size() int {
	cq.lock.Lock()
	defer cq.lock.Unlock()
	return cq.count
}
