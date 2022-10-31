package problems

func maxSubArray(nums []int) int {
	result, cur := nums[0], 0
	for _, num := range nums {
		cur += num
		if cur > result {
			result = cur
		}
		if cur < 0 {
			cur = 0
		}
	}

	return result
}
