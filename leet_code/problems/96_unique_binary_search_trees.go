package problems

func numTrees(n int) int {
	if n < 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := i - 1; j >= 0; j-- {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}
	return dp[n]
}
