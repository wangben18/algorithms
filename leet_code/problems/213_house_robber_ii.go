package problems

func rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(rob(nums[:len(nums)-1]), rob(nums[1:]))
}
