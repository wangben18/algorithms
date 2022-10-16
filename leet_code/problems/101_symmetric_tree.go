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
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	stack := Stack[*TreeNode]{}
	stack.Push(root.Right)
	stack.Push(root.Left)
	for !stack.Empty() {
		node1 := stack.Pop()
		node2 := stack.Pop()
		if node1 == nil && node2 == nil {
			continue
		}
		if node1 == nil || node2 == nil || (node1.Val != node2.Val) {
			return false
		}
		stack.Push(node1.Left)
		stack.Push(node2.Right)
		stack.Push(node1.Right)
		stack.Push(node2.Left)
	}

	return true
}

// 递归法
// func isSymmetric(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}
// 	var symmetric func(left *TreeNode, right *TreeNode) bool
// 	symmetric = func(left, right *TreeNode) bool {
// 		if left == nil && right != nil {
// 			return false
// 		} else if left != nil && right == nil {
// 			return false
// 		} else if left == nil && right == nil {
// 			return true
// 		} else if left.Val != right.Val {
// 			return false
// 		}
// 		return symmetric(left.Left, right.Right) && symmetric(left.Right, right.Left)
// 	}
// 	return symmetric(root.Left, root.Right)
// }
