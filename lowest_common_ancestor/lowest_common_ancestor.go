package lowest_common_ancestor

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	switch {
	case left != nil && right != nil:
		return root
	case left != nil && right == nil:
		return left
	case left == nil && right != nil:
		return right
	default:
		return nil
	}
}
