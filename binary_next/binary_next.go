package binarynext

type BinaryTreeNode struct {
	Left, Right, Parent *BinaryTreeNode
	Val                 int
}

func (n *BinaryTreeNode) Next() *BinaryTreeNode {
	if n.Right != nil {
		cur := n.Right
		for cur.Left != nil {
			cur = cur.Left
		}
		return cur
	}
	if n.Parent == nil {
		return nil
	}
	if n == n.Parent.Left {
		return n.Parent
	} else {
		cur := n.Parent
		for cur.Parent != nil && cur == cur.Parent.Right {
			cur = cur.Parent
		}
		if cur.Parent == nil {
			return nil
		}
		return cur.Parent
	}
}
