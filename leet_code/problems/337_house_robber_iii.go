package problems

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob3(root *TreeNode) int {
	return max(robTree(root))
}

func robTree(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	left0, left1 := robTree(root.Left)
	right0, right1 := robTree(root.Right)
	val0 := max(left0, left1) + max(right0, right1)
	val1 := root.Val + left0 + right0
	return val0, val1
}
