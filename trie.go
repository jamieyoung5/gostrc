package gostrc

import "sync"

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
	lock sync.RWMutex
}

func NewTrie(elems []string) *Trie {
	trie := &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
			isEnd:    false,
		},
	}

	for _, elem := range elems {
		trie.Insert(elem)
	}

	return trie
}

func (t *Trie) Insert(word string) {
	t.lock.Lock()
	defer t.lock.Unlock()

	if t.root == nil {
		t.root = &TrieNode{children: make(map[rune]*TrieNode)}
	}

	currentNode := t.root
	for _, char := range word {
		if _, ok := currentNode.children[char]; !ok {
			currentNode.children[char] = &TrieNode{
				children: make(map[rune]*TrieNode),
				isEnd:    false,
			}
		}

		currentNode = currentNode.children[char]
	}
	currentNode.isEnd = true
}

func (t *Trie) SearchNode(subject string) (*TrieNode, bool) {
	t.lock.RLock()
	defer t.lock.RUnlock()

	currentNode := t.root
	for _, char := range subject {
		if _, exists := currentNode.children[char]; !exists {
			return nil, false
		}
		currentNode = currentNode.children[char]
	}

	return currentNode, true
}

func (t *Trie) Search(word string) bool {
	node, ok := t.SearchNode(word)
	return ok && node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	_, ok := t.SearchNode(prefix)
	return ok
}
