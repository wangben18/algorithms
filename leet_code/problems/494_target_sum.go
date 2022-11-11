package problems

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if (sum+target)%2 == 1 {
		return 0
	}
	if target > sum || -target > sum {
		return 0
	}
	left := (sum + target) / 2
	dp := make([]int, left+1)
	dp[0] = 1
	for _, num := range nums {
		for i := left; i >= num; i-- {
			dp[i] += dp[i-num]
		}
	}
	return dp[left]
}
