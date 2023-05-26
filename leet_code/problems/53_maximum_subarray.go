package problems

// 动态规划
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	result := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		if dp[i] > result {
			result = dp[i]
		}
	}

	return result
}

// 贪心
// func maxSubArray(nums []int) int {
// 	result, cur := nums[0], 0
// 	for _, num := range nums {
// 		cur += num
// 		if cur > result {
// 			result = cur
// 		}
// 		if cur < 0 {
// 			cur = 0
// 		}
// 	}

// 	return result
// }
