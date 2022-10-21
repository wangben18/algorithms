package problems

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	queue := TreeNodeQueue{}
	queue.Push(root)
	for !queue.Empty() {
		levelMax := 0
		length := queue.Size()
		for i := 0; i < length; i++ {
			node := queue.Pop()
			if node.Left != nil {
				queue.Push(node.Left)
			}
			if node.Right != nil {
				queue.Push(node.Right)
			}
			if i == 0 || node.Val > levelMax {
				levelMax = node.Val
			}
		}
		result = append(result, levelMax)
	}
	return result
}
