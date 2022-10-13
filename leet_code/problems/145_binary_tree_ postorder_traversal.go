package problems

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	var result []int
	var traversal func(*TreeNode)
	traversal = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		traversal(cur.Left)
		traversal(cur.Right)
		result = append(result, cur.Val)
	}
	traversal(root)
	return result
}
