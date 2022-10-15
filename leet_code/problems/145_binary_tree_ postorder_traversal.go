package problems

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 迭代法
func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	stack := Stack[*TreeNode]{}
	stack.Push(root)
	for !stack.Empty() {
		node := stack.Pop()
		if node == nil {
			continue
		}
		result = append(result, node.Val)
		stack.Push(node.Left)
		stack.Push(node.Right)
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

// 递归法
// func postorderTraversal(root *TreeNode) []int {
// 	var result []int
// 	var traversal func(*TreeNode)
// 	traversal = func(cur *TreeNode) {
// 		if cur == nil {
// 			return
// 		}
// 		traversal(cur.Left)
// 		traversal(cur.Right)
// 		result = append(result, cur.Val)
// 	}
// 	traversal(root)
// 	return result
// }
