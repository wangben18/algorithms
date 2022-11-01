package problems

import "sort"

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	result := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 && k > 0 {
			nums[i] = -nums[i]
			k -= 1
		}
		result += nums[i]
	}
	if k > 0 {
		sort.Ints(nums)
		if k%2 == 1 {
			nums[0] = -nums[0]
			result -= 2 * nums[0]
		}
	}
	return result
}
