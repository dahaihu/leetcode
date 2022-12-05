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
	end := node
	{
		leftStart, leftEnd := doFlatten(left)
		if leftStart != nil {
			end.Right = leftStart
			end = leftEnd
		}
	}
	{
		rightStart, rightEnd := doFlatten(right)
		if rightStart != nil {
			end.Right = rightStart
			end = rightEnd
		}
	}
	return node, end
}
