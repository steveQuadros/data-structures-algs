package main

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.Insert("testing")
	require.True(t, trie.Search("testing"))
	require.False(t, trie.Search("test"))
	require.True(t, trie.StartsWith("test"))
	require.False(t, trie.StartsWith("ing"))
}