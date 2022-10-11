package lru

import "fmt"

type LRU[K comparable, V any] struct {
	Values               map[K]*Node[K, V]
	DummyHead, DummyTail *Node[K, V]
	Capacity             int
}

func NewLRU[K comparable, V any](capacity int) *LRU[K, V] {
	if capacity <= 0 {
		panic(fmt.Errorf("invalid capacity %d", capacity))
	}
	var (
		key   K
		value V
	)
	head := NewNode[K, V](key, value)
	tail := NewNode[K, V](key, value)
	head.Next, tail.Prev = tail, head
	return &LRU[K, V]{
		Values:    make(map[K]*Node[K, V], capacity),
		DummyHead: head,
		DummyTail: tail,
		Capacity:  capacity,
	}
}

func (l *LRU[K, V]) moveToHead(n *Node[K, V]) {
	l.listRemove(n)
	l.listAdd(n)
}

func (l *LRU[K, V]) add(key K, value V) {
	node := NewNode[K, V](key, value)
	l.Values[key] = node
	l.listAdd(node)
}

func (l *LRU[K, V]) listRemove(n *Node[K, V]) {
	n.Prev.Next, n.Next.Prev = n.Next, n.Prev
}

func (l *LRU[K, V]) listAdd(n *Node[K, V]) {
	// move to head
	oldHead := l.DummyHead.Next
	// dummyHead <-> n
	l.DummyHead.Next, n.Prev = n, l.DummyHead
	// n <-> oldHead
	n.Next, oldHead.Prev = oldHead, n

}

func (l *LRU[K, V]) evict() {
	tail := l.DummyTail.Prev
	delete(l.Values, tail.Key)
	tail.Prev.Next, tail.Next.Prev = tail.Next, tail.Prev
}

func (l *LRU[K, V]) Set(key K, value V) {
	node, existed := l.Values[key]
	if existed {
		node.Value = value
		l.moveToHead(node)
		return
	}
	if len(l.Values) == l.Capacity {
		l.evict()
	}
	l.add(key, value)
}

func (l *LRU[K, V]) Get(key K) (v V) {
	node, existed := l.Values[key]
	if !existed {
		return v
	}
	l.moveToHead(node)
	return node.Value
}

func (l *LRU[K, V]) Len() int {
	return len(l.Values)
}
