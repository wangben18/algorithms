package problems

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if getHeight(root) == -1 {
		return false
	} else {
		return true
	}
}

func getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftHeight := getHeight(node.Left)
	if leftHeight == -1 {
		return -1
	}
	rightHeight := getHeight(node.Right)
	if rightHeight == -1 {
		return -1
	}
	if int(math.Abs(float64(leftHeight)-float64(rightHeight))) > 1 {
		return -1
	} else {
		if leftHeight > rightHeight {
			return leftHeight + 1
		} else {
			return rightHeight + 1
		}
	}
}
