package problems

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	var pre *TreeNode
	maxCount, count := 1, 1
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		if pre != nil && pre.Val == node.Val {
			count++
		} else {
			count = 1
		}
		if count == maxCount {
			result = append(result, node.Val)
		} else if count > maxCount {
			maxCount = count
			result = []int{node.Val}
		}

		traversal(node.Right)
	}
	traversal(root)
	return result
}
