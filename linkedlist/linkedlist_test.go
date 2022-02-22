package linkedlist

import (
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestDoublyLinkedList(t *testing.T) {
	ddl := NewDoublyLinkedList()

	one := NewDoublyLinkedListNode("1")
	ddl.PushBack(&one)
	require.Equal(t, &one, ddl.Front())
	require.Equal(t, &one, ddl.Back())

	two := NewDoublyLinkedListNode("2")
	ddl.PushBack(&two)
	require.Equal(t, &one, ddl.Front())
	require.Equal(t, &two, ddl.Back())

	three := NewDoublyLinkedListNode("3")
	ddl.PushBack(&three)
	require.Equal(t, &one, ddl.Front())
	require.Equal(t, &three, ddl.Back())

	// remove head
	ddl.Remove(&one)
	require.Equal(t, &two, ddl.Front())
	require.Equal(t, &three, ddl.Back())

	// remove tail
	ddl.Remove(&two)
	require.Equal(t, &three, ddl.Front())
	require.Equal(t, &three, ddl.Back())

	// pushFront
	four := NewDoublyLinkedListNode("4")
	ddl.PushFront(&four)
	require.Equal(t, &four, ddl.Front())
	require.Equal(t, &three, ddl.Back())

	// pushFront
	five := NewDoublyLinkedListNode("5")
	ddl.PushFront(&five)
	require.Equal(t, &five, ddl.Front())
	require.Equal(t, &three, ddl.Back())
}

func TestDoublyLinkedList_PushFront(t *testing.T) {
	ddl := NewDoublyLinkedList()
	var nodes []DoublyLinkedListNode
	for i := 0; i < 10; i++ {
		n := NewDoublyLinkedListNode(strconv.Itoa(i))
		nodes = append(nodes, n)
		ddl.PushFront(&n)
	}

	require.Equal(t, nodes[9].Val, ddl.Front().Val, ddl.String())
	require.Equal(t, nodes[0].Val, ddl.Back().Val, ddl.String())
}

func TestDoublyLinkedList_PushBack(t *testing.T) {
	ddl := NewDoublyLinkedList()
	var nodes []DoublyLinkedListNode
	for i := 0; i < 10; i++ {
		n := NewDoublyLinkedListNode(strconv.Itoa(i))
		nodes = append(nodes, n)
		ddl.PushBack(&n)
	}

	require.Equal(t, nodes[0].Val, ddl.Front().Val, ddl.String())
	require.Equal(t, nodes[9].Val, ddl.Back().Val, ddl.String())
}

func TestDoublyLinkedList_IsEmpty(t *testing.T) {
	ddl := NewDoublyLinkedList()
	require.True(t, ddl.IsEmpty())

	node := NewDoublyLinkedListNode("1")
	ddl.PushBack(&node)
	require.False(t, ddl.IsEmpty())

	ddl.Remove(&node)
	require.True(t, ddl.IsEmpty())
}
