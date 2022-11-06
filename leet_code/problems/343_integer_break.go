package problems

func integerBreak(n int) int {
	dp := []int{0, 0, 1, 2, 4, 6, 9}
	for i := 7; i <= n; i++ {
		dp = append(dp, dp[i-3]*3)
	}
	return dp[n]
}
