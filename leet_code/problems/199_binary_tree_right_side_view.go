package problems

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	order := levelOrder(root)
	result := make([]int, 0)
	for _, cur := range order {
		result = append(result, cur[len(cur)-1])
	}
	return result
}
