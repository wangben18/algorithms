package problems

// 迭代法（统一实现）
func preorderTraversal(root *TreeNode) []int {
	var result []int
	stack := Stack[*TreeNode]{}
	if root == nil {
		return nil
	}
	stack.Push(root)
	for !stack.Empty() {
		node := stack.Pop()
		if node == nil {
			result = append(result, stack.Pop().Val)
			continue
		}
		if node.Right != nil {
			stack.Push(node.Right)
		}
		if node.Left != nil {
			stack.Push(node.Left)
		}
		stack.Push(node)
		stack.Push(nil)
	}
	return result
}

// 迭代法
// func preorderTraversal(root *TreeNode) []int {
// 	var result []int
// 	stack := Stack[*TreeNode]{}
// 	if root == nil {
// 		return nil
// 	}
// 	stack.Push(root)
// 	for !stack.Empty() {
// 		node := stack.Pop()
// 		result = append(result, node.Val)
// 		if node.Right != nil {
// 			stack.Push(node.Right)
// 		}
// 		if node.Left != nil {
// 			stack.Push(node.Left)
// 		}
// 	}
// 	return result
// }

// 递归法
// func preorderTraversal(root *TreeNode) []int {
// 	var result []int
// 	var traversal func(*TreeNode)
// 	traversal = func(cur *TreeNode) {
// 		if cur == nil {
// 			return
// 		}
// 		result = append(result, cur.Val)
// 		traversal(cur.Left)
// 		traversal(cur.Right)
// 	}
// 	traversal(root)
// 	return result
// }
