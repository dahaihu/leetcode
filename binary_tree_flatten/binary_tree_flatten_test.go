package binarytreeflatten

import (
	"fmt"
	"testing"
)

func Test_flatten(t *testing.T) {
	n := &TreeNode{Val: 1}
	left := &TreeNode{Val: 2}
	right := &TreeNode{Val: 3}
	n.Left, n.Right = left, right

	_, _ = doFlatten(n)
	for n != nil {
		fmt.Printf("->%d", n.Val)
		n = n.Right
	}
}
