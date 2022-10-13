package problems

func preorderTraversal(root *TreeNode) []int {
	var result []int
	var traversal func(*TreeNode)
	traversal = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		result = append(result, cur.Val)
		traversal(cur.Left)
		traversal(cur.Right)
	}
	traversal(root)
	return result
}
