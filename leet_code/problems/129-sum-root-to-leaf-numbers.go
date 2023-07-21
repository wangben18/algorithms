package problems

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	sum := 0
	path := make([]int, 0)

	pathToInt := func(path []int) (num int) {
		for _, v := range path {
			num = num*10 + v
		}
		return num
	}

	var traverse func(tree *TreeNode)
	traverse = func(tree *TreeNode) {
		if tree == nil {
			return
		}
		path = append(path, tree.Val)
		traverse(tree.Left)
		if tree.Left == nil && tree.Right == nil {
			sum += pathToInt(path)
		}
		traverse(tree.Right)
		path = path[:len(path)-1]
	}

	traverse(root)

	return sum
}
