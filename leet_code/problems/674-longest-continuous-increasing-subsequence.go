package problems

func findLengthOfLCIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	res := 1
	for i := 1; i < len(dp); i++ {
		if nums[i-1] < nums[i] {
			dp[i] = dp[i-1] + 1
			if dp[i] > res {
				res = dp[i]
			}
		}
	}
	return res
}
