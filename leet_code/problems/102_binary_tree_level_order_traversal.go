package problems

// 递归法
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var order func(node *TreeNode, depth int)
	result := make([][]int, 0)
	order = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == len(result) {
			result = append(result, []int{})
		}
		result[depth] = append(result[depth], node.Val)
		order(node.Left, depth+1)
		order(node.Right, depth+1)
	}
	order(root, 0)
	return result
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 迭代法
// func levelOrder(root *TreeNode) [][]int {
// 	if root == nil {
// 		return nil
// 	}
// 	queue := Queue[*TreeNode]{}
// 	queue.Push(root)
// 	result := make([][]int, 0)
// 	levelResult := make([]int, 0)
// 	for !queue.Empty() {
// 		length := queue.Size()
// 		for i := 0; i < length; i++ {
// 			node := queue.Pop()
// 			levelResult = append(levelResult, node.Val)
// 			if node.Left != nil {
// 				queue.Push(node.Left)
// 			}
// 			if node.Right != nil {
// 				queue.Push(node.Right)
// 			}
// 		}
// 		result = append(result, levelResult)
// 		levelResult = []int{}
// 	}
// 	return result
// }
