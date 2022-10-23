package problems

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return traversal(root, targetSum-root.Val)
}

func traversal(node *TreeNode, count int) bool {
	if node.Left == nil && node.Right == nil && count == 0 {
		return true
	}
	if node.Left == nil && node.Right == nil {
		return false
	}
	if node.Left != nil {
		if traversal(node.Left, count-node.Left.Val) {
			return true
		}
	}
	if node.Right != nil {
		if traversal(node.Right, count-node.Right.Val) {
			return true
		}
	}
	return false
}
