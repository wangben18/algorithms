package problems

func uniquePaths(m int, n int) int {
	dp := make([][]int, n)
	for i, _ := range dp {
		dp[i] = make([]int, m)
	}
	for row := 0; row < n; row++ {
		for col := 0; col < m; col++ {
			if row == 0 || col == 0 {
				dp[row][col] = 1
			}
			dp[row][col] = dp[row-1][col] + dp[row][col-1]
		}
	}
	return dp[n-1][m-1]
}
