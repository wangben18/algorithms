package problems

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}
	min := 0
	if root.Left != nil {
		min = root.Val - root.Left.Val
	} else if root.Right != nil {
		min = root.Right.Val - root.Val
	}
	var pre *TreeNode
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		if pre != nil && node.Val-pre.Val < min {
			min = node.Val - pre.Val
		}
		pre = node
		traversal(node.Right)
	}
	traversal(root)
	return min
}
