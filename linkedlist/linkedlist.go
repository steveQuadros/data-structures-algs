package linkedlist

import "strings"

type DoublyLinkedListNode struct {
	Prev, Next *DoublyLinkedListNode
	Val        string
}

func NewDoublyLinkedListNode(s string) DoublyLinkedListNode {
	return DoublyLinkedListNode{Val: s}
}

type DoublyLinkedList struct {
	dummyHead, dummyTail *DoublyLinkedListNode
}

func NewDoublyLinkedList() DoublyLinkedList {
	head := &DoublyLinkedListNode{}
	tail := &DoublyLinkedListNode{}
	head.Next = tail
	tail.Prev = head
	return DoublyLinkedList{
		dummyHead: head,
		dummyTail: tail,
	}
}

func (d *DoublyLinkedList) PushFront(node *DoublyLinkedListNode) {
	oldHead := d.dummyHead.Next
	d.dummyHead.Next = node
	node.Next = oldHead
	node.Prev = d.dummyHead
	oldHead.Prev = node
}

func (d *DoublyLinkedList) PushBack(node *DoublyLinkedListNode) {
	oldTail := d.dummyTail.Prev
	node.Next = d.dummyTail
	node.Prev = oldTail
	oldTail.Next = node
	d.dummyTail.Prev = node
}

func (d *DoublyLinkedList) Remove(node *DoublyLinkedListNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	node = nil
}

func (d *DoublyLinkedList) IsEmpty() bool {
	return d.dummyHead.Next == d.dummyTail
}

func (d *DoublyLinkedList) Front() *DoublyLinkedListNode {
	return d.dummyHead.Next
}

func (d *DoublyLinkedList) Back() *DoublyLinkedListNode {
	return d.dummyTail.Prev
}

func (d *DoublyLinkedList) String() string {
	node := d.dummyHead.Next
	b := strings.Builder{}
	for node != nil {
		b.WriteString(node.Val)
		node = node.Next
		if node != nil {
			b.WriteString(" <-> ")
		}
	}
	return b.String()
}
