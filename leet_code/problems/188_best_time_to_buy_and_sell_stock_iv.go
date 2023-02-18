package problems

// dp[2*i]: 第 i+1 次持有
// dp[2*i+1]: 第 i+1 次卖出
func maxProfit4(k int, prices []int) int {
	dp := make([]int, 2*k)
	for i := 0; i < len(dp); i += 2 {
		dp[i] = -prices[0]
	}
	for i := 1; i < len(prices); i++ {
		for j, pre := 0, 0; j < len(dp); pre, j = dp[j+1], j+2 {
			dp[j] = max(dp[j], pre-prices[i])
			dp[j+1] = max(dp[j+1], dp[j]+prices[i])
		}
	}
	return dp[len(dp)-1]
}
