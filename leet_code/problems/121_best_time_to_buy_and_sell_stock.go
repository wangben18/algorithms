package problems

import "math"

// 贪心
func maxProfit(prices []int) int {
	min := math.MaxInt32
	result := 0
	for _, price := range prices {
		if price < min {
			min = price
		}
		if price-min > result {
			result = price - min
		}
	}
	return result
}
