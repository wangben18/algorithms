package problems

// 动态规划（滚动数组）
// func maxProfit(prices []int) int {
// 	dp := [2][2]int{}
// 	dp[0][0] = -prices[0]
// 	for i := 1; i < len(prices); i++ {
// 		dp[i%2][0] = max(dp[(i+1)%2][0], -prices[i])
// 		dp[i%2][1] = max(dp[(i+1)%2][1], prices[i]+dp[(i+1)%2][0])
// 	}
// 	return dp[(len(prices)+1)%2][1]
// }

// 动态规划
// dp[i][0]：第 i 天持有股票所得最大现金
// dp[i][1]：第 i 天不持有股票所得最大现金
func maxProfit(prices []int) int {
	dp := make([][2]int, len(prices))
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[len(prices)-1][1]
}

// 贪心
// func maxProfit(prices []int) int {
// 	min := math.MaxInt32
// 	result := 0
// 	for _, price := range prices {
// 		if price < min {
// 			min = price
// 		}
// 		if price-min > result {
// 			result = price - min
// 		}
// 	}
// 	return result
// }
