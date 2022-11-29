package reorderlist

import (
	"bytes"
	"fmt"
)

type Node[T any] struct {
	Val  T
	Next *Node[T]
}

func (n *Node[T]) String() string {
	b := bytes.Buffer{}
	cur := n
	for cur != nil {
		b.WriteString(fmt.Sprintf("->%v", cur.Val))
		cur = cur.Next
	}
	return b.String()
}

func reverseList[T any](n *Node[T]) *Node[T] {
	var pre *Node[T]
	for n != nil {
		next := n.Next
		n.Next = pre
		pre = n
		n = next
	}
	return pre
}

func reorderList[T any](head *Node[T]) *Node[T] {
	if head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	next := slow.Next
	slow.Next = nil

	first, second := head, reverseList(next)

	dummy := &Node[T]{}
	cur := dummy
	for second != nil {
		cur.Next = first
		cur = first
		first = first.Next

		cur.Next = second
		cur = second
		second = second.Next
	}
	if first != nil {
		cur.Next = first
	}
	return dummy.Next
}
