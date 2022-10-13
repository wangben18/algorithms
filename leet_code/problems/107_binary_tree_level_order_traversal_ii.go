package problems

// 递归法
func levelOrderBottom(root *TreeNode) [][]int {
	result := levelOrder(root)
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}
