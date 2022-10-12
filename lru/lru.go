package lru

import (
	"errors"
	"fmt"
)

type LRU[K comparable, V any] struct {
	Values   map[K]*Node[K, V]
	List     *List[K, V]
	Capacity int
}

func NewLRU[K comparable, V any](capacity int) *LRU[K, V] {
	if capacity <= 0 {
		panic(fmt.Errorf("invalid capacity %d", capacity))
	}

	return &LRU[K, V]{
		Values:   make(map[K]*Node[K, V], capacity),
		List:     NewList[K, V](),
		Capacity: capacity,
	}
}

func (l *LRU[K, V]) add(key K, value V) {
	node := l.List.PushHead(key, value)
	l.Values[key] = node
}

func (l *LRU[K, V]) evict() {
	tail := l.List.Tail()
	if tail != nil {
		l.List.Remove(tail)
		delete(l.Values, tail.Key)
	}
}

func (l *LRU[K, V]) Set(key K, value V) {
	node, existed := l.Values[key]
	if existed {
		node.Value = value
		l.List.MoveToHead(node)
		return
	}
	if len(l.Values) == l.Capacity {
		l.evict()
	}
	l.add(key, value)
}

var ErrorNotExisted = errors.New("key not exist")

func (l *LRU[K, V]) Get(key K) (v V, err error) {
	node, existed := l.Values[key]
	if !existed {
		return v, ErrorNotExisted
	}
	l.List.MoveToHead(node)
	return node.Value, nil
}

func (l *LRU[K, V]) Len() int {
	return len(l.Values)
}
