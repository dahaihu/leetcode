package lru

import (
	"bytes"
	"fmt"
)

type List[K comparable, V any] struct {
	DummyHead, DummyTail *Node[K, V]
}

func NewList[K comparable, V any]() *List[K, V] {
	l := new(List[K, V])
	var (
		key   K
		value V
	)
	head := NewNode[K, V](key, value, l)
	tail := NewNode[K, V](key, value, l)
	head.Next, tail.Prev = tail, head
	l.DummyHead, l.DummyTail = head, tail
	return l
}

func (l *List[K, V]) String() string {
	buf := bytes.Buffer{}
	cur := l.DummyHead.Next
	for cur != l.DummyTail {
		if cur.Next == l.DummyTail {
			buf.WriteString(fmt.Sprintf("Node(%v, %v)", cur.Key, cur.Value))
		} else {
			buf.WriteString(fmt.Sprintf("Node(%v, %v)<->", cur.Key, cur.Value))
		}
		cur = cur.Next
	}
	return buf.String()
}

func (l *List[K, V]) Remove(n *Node[K, V]) (removed bool) {
	if n.List != l {
		return false
	}
	n.Prev.Next, n.Next.Prev = n.Next, n.Prev

	n.Prev, n.Next = nil, nil
	return true
}

func (l *List[K, V]) PushHead(key K, value V) *Node[K, V] {
	node := NewNode[K, V](key, value, l)
	_ = l.MoveToHead(node)
	return node
}

func (l *List[K, V]) PushTail(key K, value V) {
	node := NewNode[K, V](key, value, l)
	_ = l.MoveToTail(node)
}

func (l *List[K, V]) Head() *Node[K, V] {
	head := l.DummyHead.Next
	if head == l.DummyTail {
		return nil
	}
	return head
}

func (l *List[K, V]) Tail() *Node[K, V] {
	tail := l.DummyTail.Prev
	if tail == l.DummyHead {
		return nil
	}
	return tail
}

func (l *List[K, V]) contains(n *Node[K, V]) bool {
	return n.List == l
}

func (l *List[K, V]) MoveToHead(node *Node[K, V]) bool {
	if !l.contains(node) {
		return false
	}
	// move to head
	oldHead := l.DummyHead.Next
	// node <-> DummyTail
	l.DummyHead.Next, node.Prev = node, l.DummyHead
	// oldNode <-> node
	node.Next, oldHead.Prev = oldHead, node
	return true
}

func (l *List[K, V]) MoveToTail(node *Node[K, V]) bool {
	if !l.contains(node) {
		return false
	}
	// move to head
	oldTail := l.DummyTail.Prev
	// node <-> DummyTail
	node.Next, l.DummyTail.Prev = l.DummyTail, node
	// oldNode <-> node
	oldTail.Next, node.Prev = node, oldTail
	return true
}
