package rb_tree

import (
	"fmt"
	"strings"
)

type color int8

const (
	black color = 0
	red   color = 1
)

type RBNode struct {
	color       color
	Key         interface{}
	Value       interface{}
	Parent      *RBNode
	Left, Right *RBNode
}

func (n *RBNode) String() string {
	return fmt.Sprintf("<%v:%v:%v>", n.Key, n.Value, n.color)
}

func (n *RBNode) Len() int {
	if n == nil {
		return 0
	}
	return len(fmt.Sprintf("<%v:%v:%v>", n.Key, n.Value, n.color))
}

func (n *RBNode) IsNil() bool {
	return n == nil
}

func (n *RBNode) LeftChild() INode {
	return n.Left
}

func (n *RBNode) RightChild() INode {
	return n.Right
}

func (n *RBNode) sibling() *RBNode {
	if n == nil || n.Parent == nil {
		return nil
	}
	if n == n.Parent.Left {
		return n.Parent.Right
	}
	return n.Parent.Left
}

func (n *RBNode) uncle() *RBNode {
	if n != nil && n.Parent != nil && n.Parent.Parent != nil {
		return n.Parent.sibling()
	}
	return nil
}

func (n *RBNode) grandparent() *RBNode {
	if n != nil && n.Parent != nil && n.Parent.Parent != nil {
		return n.Parent.Parent
	}
	return nil
}

type RBTree struct {
	Root    *RBNode
	size    int64
	Compare func(a, b interface{}) int
}

func (t *RBTree) String() string {
	lines, _, _, _ := buildTree(t.Root)
	return strings.Join(lines, "\n")
}
func (t *RBTree) Set(key, value interface{}) {
	if t.Root == nil {
		t.Root = &RBNode{Key: key, Value: value, Parent: nil, color: black}
		t.size += 1
		return
	}
	var node *RBNode
	cur := t.Root
	for {
		if compareResult := t.Compare(key, cur.Key); compareResult > 0 {
			if cur.Right != nil {
				cur = cur.Right
			} else {
				node = &RBNode{Key: key, Value: value, color: red, Parent: cur}
				cur.Right = node
				break
			}
		} else if compareResult < 0 {
			if cur.Left != nil {
				cur = cur.Left
			} else {
				node = &RBNode{Key: key, Value: value, color: red, Parent: cur}
				cur.Left = node
				break
			}
		} else {
			cur.Value = value
			return
		}
	}
	t.insertCase1(node)
	t.size += 1
}

func (t *RBTree) insertCase1(n *RBNode) {
	if n.Parent == nil {
		n.color = black
		return
	}
	t.insertCase2(n)
}

func (t *RBTree) insertCase2(n *RBNode) {
	if n.Parent.color == black {
		return
	}
	t.insertCase3(n)
}

func (t *RBTree) insertCase3(n *RBNode) {
	if uncle := n.uncle(); uncle != nil && uncle.color == red {
		n.Parent.color = black
		uncle.color = black
		grandparent := n.grandparent()
		grandparent.color = red
		t.insertCase1(grandparent)
		return
	}
	t.insertCase4(n)
}

func (t *RBTree) rotateRight(n *RBNode) {
	left := n.Right
	n.setLeft(left.Right)
	if n.Parent != nil {
		if n == n.Parent.Left {
			n.Parent.setLeft(left)
		} else {
			n.Parent.setRight(left)
		}
	} else {
		t.setRoot(left)
	}
	left.setRight(n)
}

func (t *RBTree) setRoot(n *RBNode) {
	t.Root = n
	n.Parent = nil
}

func (t *RBTree) rotateLeft(n *RBNode) {
	right := n.Right
	n.setRight(right.Left)
	if n.Parent != nil {
		if n == n.Parent.Left {
			n.Parent.setLeft(right)
		} else {
			n.Parent.setRight(right)
		}
	} else {
		t.setRoot(right)
	}
	right.setLeft(n)
}

func (n *RBNode) setLeft(left *RBNode) {
	n.Left = left
	if left != nil {
		left.Parent = n
	}
}

func (n *RBNode) setRight(right *RBNode) {
	n.Right = right
	if right != nil {
		right.Parent = n
	}
}

func (t *RBTree) insertCase4(n *RBNode) {
	grandparent := n.grandparent()
	if n == n.Parent.Right && n.Parent == grandparent.Left {
		t.rotateLeft(n.Parent)
		n = n.Left
	} else if n == n.Parent.Left && n.Parent == grandparent.Right {
		t.rotateRight(n.Parent)
		n = n.Right
	}
	t.insertCase5(n)
}

func (t *RBTree) insertCase5(n *RBNode) {
	n.Parent.color = black
	grandparent := n.grandparent()
	grandparent.color = red
	if n == n.Parent.Left {
		t.rotateRight(grandparent)
	} else {
		t.rotateLeft(grandparent)
	}
}

func (t *RBTree) lookup(key interface{}) *RBNode {
	cur := t.Root
	for {
		compareResult := t.Compare(key, cur.Key)
		if compareResult > 0 {
			if cur.Right == nil {
				return nil
			}
			cur = cur.Right
		} else if compareResult < 0 {
			if cur.Left == nil {
				return nil
			}
			cur = cur.Left
		} else {
			return cur
		}
	}
}

func (n *RBNode) maximumNode() *RBNode {
	cur := n
	for cur.Right != nil {
		cur = cur.Right
	}
	return cur
}

func (t *RBTree) replace(old, new *RBNode) {
	if old.Parent == nil {
		t.Root = new
	} else {
		if old == old.Parent.Left {
			old.Parent.Left = new
		} else {
			old.Parent.Right = new
		}
	}
	if new != nil {
		new.Parent = old.Parent
	}
}

func (n *RBNode) getColor() color {
	if n == nil {
		return black
	}
	return n.color
}

func (n *RBNode) isNil() bool {
	return n == nil
}

func (n *RBNode) isBlack() bool {
	return n.getColor() == black
}

func (n *RBNode) isRed() bool {
	return n.getColor() == red
}

func (t *RBTree) Remove(key interface{}) {
	node := t.lookup(key)
	if node == nil {
		return
	}
	if node.Left != nil && node.Right != nil {
		replaced := node.Left.maximumNode()
		node.Key = replaced.Key
		node.Value = replaced.Value
		node = replaced
	}
	if node.Left == nil || node.Right == nil {
		var child *RBNode
		if node.Right == nil {
			child = node.Right
		} else {
			child = node.Left
		}
		if node.isBlack() {
			node.color = child.getColor()
			t.deleteCase1(node)
		}
		if t.Root == node && child != nil {
			child.color = black
		}
		t.replace(node, child)
	}
	t.size -= 1
}

func (t *RBTree) deleteCase1(n *RBNode) {
	if n.Parent == nil {
		return
	}
	t.deleteCase2(n)
}

func (t *RBTree) deleteCase2(n *RBNode) {
	sibling := n.sibling()
	if sibling.isRed() {
		n.Parent.color = red
		sibling.color = black
		if n == n.Parent.Left {
			t.rotateLeft(n.Parent)
		} else {
			t.rotateRight(n.Parent)
		}
	}
	t.deleteCase3(n)
}

func (t *RBTree) deleteCase3(n *RBNode) {
	sibling := n.sibling()
	if n.Parent.isBlack() &&
		sibling.isBlack() &&
		sibling.Left.isBlack() &&
		sibling.Right.isBlack() {
		sibling.color = red
		t.deleteCase1(n.Parent)
	} else {
		t.deleteCase4(n)
	}
}

func (t *RBTree) deleteCase4(n *RBNode) {
	if n.Parent.isRed() {
		sibling := n.sibling()
		n.Parent.color = black
		sibling.color = red
		return
	}
	t.deleteCase5(n)
}

func (t *RBTree) deleteCase5(n *RBNode) {
	sibling := n.sibling()
	if n == n.Parent.Left &&
		sibling.isBlack() &&
		sibling.Left.isRed() &&
		sibling.Right.isBlack() {
		sibling.color = red
		sibling.Left.color = black
		t.rotateRight(sibling)
	} else if n == n.Parent.Right &&
		sibling.isBlack() &&
		sibling.Right.isRed() &&
		sibling.Left.isBlack() {
		sibling.color = red
		sibling.Right.color = black
		t.rotateLeft(sibling)
	}
	t.deleteCase6(n)
}

func (t *RBTree) deleteCase6(n *RBNode) {
	sibling := n.sibling()
	if n == n.Parent.Left && sibling.isBlack() && sibling.Right.isRed() {
		sibling.Right.color = black
		t.rotateLeft(n.Parent)
	} else if n == n.Parent.Right && sibling.isBlack() && sibling.Left.isRed() {
		sibling.Left.color = black
		t.rotateRight(n.Parent)
	}
}
