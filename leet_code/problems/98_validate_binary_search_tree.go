package problems

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var pre *TreeNode
	var traverse func(node *TreeNode) bool
	traverse = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		left := traverse(node.Left)
		if pre != nil && pre.Val >= node.Val {
			return false
		}
		pre = node
		right := traverse(node.Right)
		return left && right
	}
	return traverse(root)
}
