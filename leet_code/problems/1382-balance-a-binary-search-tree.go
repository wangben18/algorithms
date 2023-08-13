package problems

func balanceBST(root *TreeNode) *TreeNode {
	var nums []int
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node.Left != nil {
			traverse(node.Left)
		}
		nums = append(nums, node.Val)
		if node.Right != nil {
			traverse(node.Right)
		}
	}
	traverse(root)
	return sortedArrayToBST(nums)
}
