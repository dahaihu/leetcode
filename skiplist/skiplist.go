package skiplist

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
)

type Node struct {
	Key   int
	Value interface{}
	Next  []*Node
}

type Skiplist struct {
	Head      *Node
	p         float64
	MaxHeight int
}

func Constructor() Skiplist {
	return Skiplist{
		Head:      new(Node),
		p:         0.5,
		MaxHeight: 10,
	}
}

func (s *Skiplist) height() int {
	h := 1
	for h < s.MaxHeight && rand.Float64() < s.p {
		h += 1
	}
	return h
}

func (s *Skiplist) prefixes(level int, key int) []*Node {
	prefixes := make([]*Node, level)
	for height, cur := level-1, s.Head; height >= 0; height -= 1 {
		for cur.Next[height] != nil && key > cur.Next[height].Key {
			cur = cur.Next[height]
		}
		prefixes[height] = cur
	}
	return prefixes
}

func (s *Skiplist) Add(key int) {
	nodeHeight := s.height()
	node := &Node{
		Key:  key,
		Next: make([]*Node, nodeHeight),
	}

	skiplistHeight := len(s.Head.Next)
	prefixLen := nodeHeight
	if nodeHeight > skiplistHeight {
		for i := skiplistHeight; i < nodeHeight; i++ {
			s.Head.Next = append(s.Head.Next, node)
		}
		prefixLen = skiplistHeight
	}

	prefixes := s.prefixes(prefixLen, key)

	for height := prefixLen - 1; height >= 0; height-- {
		prefix := prefixes[height]
		next := prefix.Next[height]
		prefix.Next[height] = node
		node.Next[height] = next
	}
}

func (s *Skiplist) Search(key int) bool {
	cur := s.Head
	for height := len(s.Head.Next) - 1; height >= 0; height-- {
		for cur.Next[height] != nil && key > cur.Next[height].Key {
			cur = cur.Next[height]
		}
		if cur.Next[height] != nil && cur.Next[height].Key == key {
			return true
		}
	}
	return false
}

func (s *Skiplist) Erase(key int) bool {
	prefixes := s.prefixes(len(s.Head.Next), key)
	var deleted bool
	for i := len(s.Head.Next) - 1; i >= 0; i-- {
		if prefixes[i] == nil {
			continue
		}
		prefix := prefixes[i]
		next := prefix.Next[i]
		if next == nil || key < next.Key {
			continue
		}
		if !deleted {
			deleted = true
		}
		prefix.Next[i] = next.Next[i]
	}
	return deleted
}

func (s *Skiplist) String() string {
	maxHeight := len(s.Head.Next)
	mark := make([][]int, maxHeight)
	var length int
	for cur := s.Head.Next[0]; cur != nil; cur = cur.Next[0] {
		length += 1
	}
	for i := 0; i < maxHeight; i++ {
		mark[i] = make([]int, length)
	}
	for cur, idx := s.Head.Next[0], 0; cur != nil; cur, idx = cur.Next[0], idx+1 {
		for level := 0; level < len(cur.Next); level++ {
			mark[level][idx] = cur.Key
		}
	}
	var buf bytes.Buffer
	for level := 0; level < maxHeight; level++ {
		buf.WriteString("head")
		var emtpyCnt int
		for idx := 0; idx < length; idx++ {
			if mark[level][idx] != 0 {
				buf.WriteString(fmt.Sprintf("%s->%03d", strings.Repeat(strings.Repeat("-", 5), emtpyCnt), mark[level][idx]))
				emtpyCnt = 0
			} else {
				emtpyCnt += 1
			}
		}
		buf.WriteString(fmt.Sprintf("%s->nil", strings.Repeat(strings.Repeat("-", 5), emtpyCnt)))
		buf.WriteString("\n")
	}
	return buf.String()
}
