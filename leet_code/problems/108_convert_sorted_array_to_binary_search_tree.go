package problems

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	var traversal func(nums []int, l, r int) *TreeNode
	traversal = func(nums []int, l, r int) *TreeNode {
		if l > r {
			return nil
		}
		mid := l + (r-l)/2
		node := &TreeNode{Val: nums[mid]}
		node.Left = traversal(nums, l, mid-1)
		node.Right = traversal(nums, mid+1, r)
		return node
	}
	return traversal(nums, 0, len(nums)-1)
}
