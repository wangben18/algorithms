package problems

func maxProfit(prices []int) int {
	result := 0
	for i := 1; i < len(prices); i++ {
		if profit := prices[i] - prices[i-1]; profit > 0 {
			result += profit
		}
	}
	return result
}
