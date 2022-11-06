package problems

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	dp := make([]int, sum/2+1)
	for i := range nums {
		for j := sum / 2; j >= nums[i]; j-- {
			if dp[j] < dp[j-nums[i]]+nums[i] {
				dp[j] = dp[j-nums[i]] + nums[i]
				if j == sum/2 && dp[j] == j {
					return true
				}
			}
		}
	}

	return false
}
