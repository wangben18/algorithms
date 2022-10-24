package problems

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	index := 0
	for i, num := range inorder {
		if num == root.Val {
			index = i
		}
	}
	inorderLeft := inorder[:index]
	inorderRight := make([]int, 0)
	if index < len(inorder) {
		inorderRight = inorder[index+1:]
	}
	root.Left = buildTree(inorderLeft, postorder[:index])
	root.Right = buildTree(inorderRight, postorder[index:len(postorder)-1])
	return root
}
