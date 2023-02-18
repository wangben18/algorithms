package problems

// 节省空间:
// 空间复杂度：O(1)
// 时间复杂度：O(n)
// dp[0]: 第一次持有
// dp[1]: 第一次卖出
// dp[2]: 第二次持有
// dp[3]: 第二次卖出
func maxProfit3(prices []int) int {
	dp := [4]int{}
	dp[0] = -prices[0]
	dp[2] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[0] = max(dp[0], -prices[i])
		dp[1] = max(dp[1], dp[0]+prices[i])
		dp[2] = max(dp[2], dp[1]-prices[i])
		dp[3] = max(dp[3], dp[2]+prices[i])
	}
	return dp[3]
}

// 空间复杂度：O(n)
// 时间复杂度：O(n)
// dp[i][0]: 无操作
// dp[i][1]: 第一次持有
// dp[i][2]: 第一次不持有
// dp[i][3]: 第二次持有
// dp[i][4]: 第二次不持有
// func maxProfit3(prices []int) int {
// 	dp := make([][5]int, len(prices))
// 	dp[0][1] = -prices[0]
// 	dp[0][3] = -prices[0]
// 	for i := 1; i < len(prices); i++ {
// 		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
// 		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
// 		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
// 		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
// 	}
// 	return dp[len(prices)-1][4]
// }
