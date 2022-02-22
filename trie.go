package main

type Trie struct {
	children map[rune]*Trie
	word bool
}

func NewTrie() Trie {
	return Trie{
		children: map[rune]*Trie{},
	}
}

func (t *Trie) Insert(word string) {
	next := t
	for _, r := range word {
		if _, ok := next.children[r]; !ok {
			trie := NewTrie()
			next.children[r] = &trie
		}
		next = next.children[r]
	}
	next.word = true
}

func(t *Trie) Search(word string) bool {
	next := t
	for _, r := range word {
		if _, ok := next.children[r]; !ok {
			return false
		}
		next = next.children[r]
	}
	return next.word
}

func (t *Trie) StartsWith(prefix string) bool {
	next := t
	for _, r := range prefix {
		if _, ok := next.children[r]; !ok {
			return false
		}
		next = next.children[r]
	}
	return true
}