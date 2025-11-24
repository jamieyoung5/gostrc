package gostrc_test

import (
	"sync"
	"testing"

	"github.com/jamieyoung5/gostrc"
)

func TestNewTrie(t *testing.T) {
	words := []string{"apple", "app", "application"}
	trie := gostrc.NewTrie(words)

	for _, word := range words {
		if !trie.Search(word) {
			t.Errorf("NewTrie failed to insert %q", word)
		}
	}

	if trie.Search("banana") {
		t.Error("NewTrie found a word that was not inserted")
	}
}

func TestTrie_InsertAndSearch(t *testing.T) {
	tests := []struct {
		name     string
		insert   []string
		search   string
		expected bool
	}{
		{"Single word match", []string{"hello"}, "hello", true},
		{"Single word mismatch", []string{"hello"}, "world", false},
		{"Prefix is not a word", []string{"hello"}, "hell", false},
		{"Word is not prefix", []string{"hell"}, "hello", false},
		{"Empty string", []string{""}, "", true},
		{"Multiple words", []string{"cat", "car", "cart"}, "car", true},
		{"Missing from multiple", []string{"cat", "car", "cart"}, "can", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trie := gostrc.NewTrie(nil)
			for _, w := range tt.insert {
				trie.Insert(w)
			}

			if got := trie.Search(tt.search); got != tt.expected {
				t.Errorf("Search(%q) = %v; want %v", tt.search, got, tt.expected)
			}
		})
	}
}

func TestTrie_StartsWith(t *testing.T) {
	tests := []struct {
		name     string
		insert   []string
		prefix   string
		expected bool
	}{
		{"Full match is prefix", []string{"hello"}, "hello", true},
		{"Partial match", []string{"hello"}, "hel", true},
		{"Single char match", []string{"hello"}, "h", true},
		{"No match", []string{"hello"}, "x", false},
		{"Superstring is not prefix", []string{"hell"}, "hello", false},
		{"Empty prefix", []string{"hello"}, "", true}, // Technically root exists
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trie := gostrc.NewTrie(nil)
			for _, w := range tt.insert {
				trie.Insert(w)
			}

			if got := trie.StartsWith(tt.prefix); got != tt.expected {
				t.Errorf("StartsWith(%q) = %v; want %v", tt.prefix, got, tt.expected)
			}
		})
	}
}

func TestTrie_SearchNode(t *testing.T) {
	trie := gostrc.NewTrie([]string{"test"})

	node, ok := trie.SearchNode("te")
	if !ok {
		t.Error("SearchNode('te') expected ok=true")
	}
	if node == nil {
		t.Error("SearchNode('te') returned nil node")
	}

	node, ok = trie.SearchNode("ta")
	if ok {
		t.Error("SearchNode('ta') expected ok=false")
	}
	if node != nil {
		t.Error("SearchNode('ta') expected nil node")
	}
}

func TestTrie_Concurrency(t *testing.T) {
	trie := gostrc.NewTrie(nil)
	var wg sync.WaitGroup

	writers := 10
	readers := 20

	for i := range writers {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			word := string(rune('a' + val))
			trie.Insert(word)
		}(i)
	}

	// Concurrent Reads
	for range readers {
		wg.Go(func() {
			trie.Search("some_word")
			trie.StartsWith("some_prefix")
		})
	}

	wg.Wait()
}
