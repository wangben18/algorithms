package problems

func inorderTraversal(root *TreeNode) []int {
	var result []int
	var traversal func(*TreeNode)
	traversal = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		traversal(cur.Left)
		result = append(result, cur.Val)
		traversal(cur.Right)
	}
	traversal(root)
	return result
}
