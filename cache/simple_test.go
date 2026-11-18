package cache

import (
	"sync"
	"testing"
)

func newTestCache() *SimpleCache[string, int] {
	return &SimpleCache[string, int]{data: map[string]int{}}
}

func TestSimpleCacheSetGetRemove(t *testing.T) {
	sc := newTestCache()

	if _, ok := sc.Get("a"); ok {
		t.Fatal("expected miss on empty cache")
	}

	sc.Set("a", 1)
	if v, ok := sc.Get("a"); !ok || v != 1 {
		t.Fatalf("Get(a) = %d, %v; want 1, true", v, ok)
	}

	sc.Set("a", 2)
	if v, _ := sc.Get("a"); v != 2 {
		t.Fatalf("Get(a) after overwrite = %d; want 2", v)
	}

	sc.Remove("a")
	if _, ok := sc.Get("a"); ok {
		t.Fatal("expected miss after Remove")
	}

	sc.Remove("missing")
}

func TestSimpleCacheConcurrent(t *testing.T) {
	sc := newTestCache()

	var wg sync.WaitGroup
	for i := range 50 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			key := string(rune('a' + i%26))
			sc.Set(key, i)
			sc.Get(key)
			sc.Remove(key)
		}()
	}
	wg.Wait()
}
