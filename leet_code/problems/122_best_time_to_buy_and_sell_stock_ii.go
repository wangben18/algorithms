package problems

// 动态规划
// dp[i][0]：第 i 天持有股票所得最大现金
// dp[i][1]：第 i 天不持有股票所得最大现金
func maxProfit2(prices []int) int {
	dp := make([][2]int, len(prices))
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i]) // 继续持有前一天的或者当天买入
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]) // 继续不持有或卖出持有的
	}
	return dp[len(prices)-1][1]
}

// 贪心
// func maxProfit2(prices []int) int {
// 	result := 0
// 	for i := 1; i < len(prices); i++ {
// 		if profit := prices[i] - prices[i-1]; profit > 0 {
// 			result += profit
// 		}
// 	}
// 	return result
// }
