package cache

import (
	"maps"
	"slices"
	"sync"
)

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

func (sc *Simple[K, V]) Values() []V {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	return slices.Collect(maps.Values(sc.data))
}

func (sc *Simple[K, V]) Len() int {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	return len(sc.data)
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

	_ = sc.initialised()

	sc.data[key] = val
}

func (sc *Simple[K, V]) SetIfAbsent(key K, val V) bool {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	if !sc.initialised() {
		sc.data[key] = val

		return true
	}

	if _, ok := sc.data[key]; ok {
		return false
	}

	sc.data[key] = val
	return true
}

func (sc *Simple[K, V]) Remove(key K) {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	delete(sc.data, key)
}

func (sc *Simple[K, V]) RemoveIf(key K, pred func(V) bool) bool {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	v, ok := sc.data[key]
	if !ok || !pred(v) {
		return false
	}

	delete(sc.data, key)

	return true
}

func (sc *Simple[K, V]) initialised() bool {
	if sc.data == nil {
		sc.data = make(map[K]V)
		return false
	}

	return true
}
