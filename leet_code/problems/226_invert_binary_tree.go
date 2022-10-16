package problems

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 层序遍历
func invertTree(root *TreeNode) *TreeNode {
	queue := Queue[*TreeNode]{}
	queue.Push(root)
	for !queue.Empty() {
		length := queue.Size()
		for i := 0; i < length; i++ {
			node := queue.Pop()
			if node.Left != nil {
				queue.Push(node.Left)
			}
			if node.Right != nil {
				queue.Push(node.Right)
			}
			node.Left, node.Right = node.Right, node.Left
		}
	}
	return root
}

// 统一迭代法
// func invertTree(root *TreeNode) *TreeNode {
// 	stack := Stack[*TreeNode]{}
// 	stack.Push(root)
// 	for !stack.Empty() {
// 		node := stack.Pop()
// 		if node == nil {
// 			node = stack.Pop()
// 			node.Left, node.Right = node.Right, node.Left
// 		} else {
// 			if node.Right != nil {
// 				stack.Push(node.Right)
// 			}
// 			if node.Left != nil {
// 				stack.Push(node.Left)
// 			}
// 			stack.Push(node)
// 			stack.Push(nil)
// 		}
// 	}
// 	return root
// }

// 迭代法
// func invertTree(root *TreeNode) *TreeNode {
// 	stack := Stack[*TreeNode]{}
// 	stack.Push(root)
// 	for !stack.Empty() {
// 		node := stack.Pop()
// 		if node == nil {
// 			continue
// 		}
// 		node.Left, node.Right = node.Right, node.Left
// 		stack.Push(node.Right)
// 		stack.Push(node.Left)
// 	}
// 	return root
// }

// 递归法
// func invertTreeRecursive(root *TreeNode) *TreeNode {
// 	if root == nil {
// 		return root
// 	}
// 	root.Left, root.Right = root.Right, root.Left
// 	invertTree(root.Left)
// 	invertTree(root.Right)
// 	return root
// }
