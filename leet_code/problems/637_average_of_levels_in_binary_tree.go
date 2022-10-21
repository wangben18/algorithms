package problems

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	queue := TreeNodeQueue{}
	result := make([]float64, 0)
	if root == nil {
		return result
	}
	queue.Push(root)
	for !queue.Empty() {
		length := queue.Size()
		curLevelSum := 0
		for i := 0; i < length; i++ {
			node := queue.Pop()
			curLevelSum += node.Val
			if node.Left != nil {
				queue.Push(node.Left)
			}
			if node.Right != nil {
				queue.Push(node.Right)
			}
		}
		result = append(result, float64(curLevelSum)/float64(length))
	}

	return result
}
