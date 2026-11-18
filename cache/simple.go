package cache

import "sync"

// SimpleCache is a map with a mutex
type SimpleCache[K comparable, V any] struct {
	mu   sync.Mutex
	data map[K]V
}

func NewSimpleCache[K comparable, V any]() *SimpleCache[K, V] {
	data := make(map[K]V)

	return &SimpleCache[K, V]{
		data: data,
	}
}

func (sc *SimpleCache[K, V]) Get(key K) (V, bool) {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	val, ok := sc.data[key]
	return val, ok
}

func (sc *SimpleCache[K, V]) Set(key K, val V) {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	if sc.data == nil {
		sc.data = make(map[K]V)
	}

	sc.data[key] = val
}

func (sc *SimpleCache[K, V]) Remove(key K) {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	delete(sc.data, key)
}
