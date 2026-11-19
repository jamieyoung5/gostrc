package cache

import (
	"slices"
	"sync"
	"testing"
)

func TestSimpleSetGetRemove(t *testing.T) {
	sc := NewSimple[string, int]()

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

func TestSimpleZeroValue(t *testing.T) {
	var sc Simple[string, int]

	if _, ok := sc.Get("a"); ok {
		t.Fatal("expected miss on zero value")
	}
	if n := sc.Len(); n != 0 {
		t.Fatalf("Len() = %d; want 0", n)
	}
	sc.Remove("a")

	sc.Set("a", 1)
	if v, ok := sc.Get("a"); !ok || v != 1 {
		t.Fatalf("Get(a) = %d, %v; want 1, true", v, ok)
	}

	var sc2 Simple[string, int]
	if !sc2.SetIfAbsent("a", 1) {
		t.Fatal("SetIfAbsent on zero value = false; want true")
	}
	if v, _ := sc2.Get("a"); v != 1 {
		t.Fatalf("Get(a) = %d; want 1", v)
	}
}

func TestSimpleSetIfAbsent(t *testing.T) {
	sc := NewSimple[string, int]()

	if !sc.SetIfAbsent("a", 1) {
		t.Fatal("SetIfAbsent on new key = false; want true")
	}
	if sc.SetIfAbsent("a", 2) {
		t.Fatal("SetIfAbsent on existing key = true; want false")
	}
	if v, _ := sc.Get("a"); v != 1 {
		t.Fatalf("Get(a) = %d; want 1", v)
	}
}

func TestSimpleRemoveIf(t *testing.T) {
	sc := NewSimple[string, int]()
	sc.Set("a", 1)

	isTwo := func(v int) bool { return v == 2 }

	if sc.RemoveIf("a", isTwo) {
		t.Fatal("RemoveIf with false predicate = true; want false")
	}
	if _, ok := sc.Get("a"); !ok {
		t.Fatal("key removed despite false predicate")
	}

	if sc.RemoveIf("missing", isTwo) {
		t.Fatal("RemoveIf on missing key = true; want false")
	}

	sc.Set("a", 2)
	if !sc.RemoveIf("a", isTwo) {
		t.Fatal("RemoveIf with true predicate = false; want true")
	}
	if _, ok := sc.Get("a"); ok {
		t.Fatal("expected miss after RemoveIf")
	}
}

func TestSimpleLenValues(t *testing.T) {
	sc := NewSimple[string, int]()
	sc.Set("a", 1)
	sc.Set("b", 2)
	sc.Set("c", 3)

	if n := sc.Len(); n != 3 {
		t.Fatalf("Len() = %d; want 3", n)
	}

	vals := sc.Values()
	slices.Sort(vals)
	if !slices.Equal(vals, []int{1, 2, 3}) {
		t.Fatalf("Values() = %v; want [1 2 3]", vals)
	}
}

func TestSimpleConcurrent(t *testing.T) {
	sc := NewSimple[string, int]()

	var wg sync.WaitGroup
	for i := range 50 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			key := string(rune('a' + i%26))
			sc.Set(key, i)
			sc.SetIfAbsent(key, i)
			sc.Get(key)
			sc.Len()
			sc.Values()
			sc.RemoveIf(key, func(int) bool { return true })
			sc.Remove(key)
		}()
	}
	wg.Wait()
}
