package binarytreeflatten

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func doFlatten(node *TreeNode) (*TreeNode, *TreeNode) {
	if node == nil {
		return nil, nil
	}
	left, right := node.Left, node.Right
	node.Left, node.Right = nil, nil
	cur := node
	// left
	{
		leftStart, leftEnd := doFlatten(left)
		if leftStart != nil {
			cur.Right = leftStart
			cur = leftEnd
		}
	}
	// right
	{
		rightStart, rightEnd := doFlatten(right)
		if rightStart != nil {
			cur.Right = rightStart
			cur = rightEnd
		}
	}
	return node, cur
}
