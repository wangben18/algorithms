package problems

// 迭代法
func inorderTraversal(root *TreeNode) []int {
	var result []int
	stack := Stack[*TreeNode]{}
	cur := root
	for cur != nil || !stack.Empty() {
		if cur != nil {
			stack.Push(cur)
			cur = cur.Left
		} else {
			cur = stack.Pop()
			result = append(result, cur.Val)
			cur = cur.Right
		}
	}
	return result
}

// 递归法
func inorderTraversalRecursive(root *TreeNode) []int {
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
