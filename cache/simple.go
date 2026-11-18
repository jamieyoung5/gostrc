package cache

import "sync"

// Simple is a map with a mutex
type Simple[K comparable, V any] struct {
	mu   sync.Mutex
	data map[K]V
}

func NewSimple[K comparable, V any]() *Simple[K, V] {
	data := make(map[K]V)

	return &Simple[K, V]{
		data: data,
	}
}

func (sc *Simple[K, V]) Get(key K) (V, bool) {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	val, ok := sc.data[key]
	return val, ok
}

func (sc *Simple[K, V]) Set(key K, val V) {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	if sc.data == nil {
		sc.data = make(map[K]V)
	}

	sc.data[key] = val
}

func (sc *Simple[K, V]) Remove(key K) {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	delete(sc.data, key)
}
