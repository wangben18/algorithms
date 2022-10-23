package problems

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	var traversal func(node *TreeNode, path string)
	result := make([]string, 0)
	traversal = func(node *TreeNode, path string) {
		path += strconv.Itoa(node.Val)
		if node.Left == nil && node.Right == nil {
			result = append(result, path)
		}
		if node.Left != nil {
			traversal(node.Left, path+"->")
		}
		if node.Right != nil {
			traversal(node.Right, path+"->")
		}
	}
	traversal(root, "")
	return result
}
