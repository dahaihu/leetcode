package lru

import (
	"bytes"
	"fmt"
)

type Node[K comparable, V any] struct {
	Key        K
	Value      V
	Next, Prev *Node[K, V]
}

func (n *Node[K, V]) String() string {
	buf := bytes.Buffer{}
	cur := n
	for cur != nil {
		if cur.Next == nil {
			buf.WriteString(fmt.Sprintf("Node(%v, %v)", cur.Key, cur.Value))
		} else {
			buf.WriteString(fmt.Sprintf("Node(%v, %v)<->", cur.Key, cur.Value))
		}
		cur = cur.Next
	}
	return buf.String()
}

func NewNode[K comparable, V any](key K, value V) *Node[K, V] {
	return &Node[K, V]{Value: value, Key: key}
}
