package problems

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, num := range stones {
		sum += num
	}
	target := sum / 2
	dp := make([]int, target+1)
	for _, stone := range stones {
		for j := target; j >= stone; j-- {
			if dp[j] < dp[j-stone]+stone {
				dp[j] = dp[j-stone] + stone
			}
		}
	}
	return sum - dp[target] - dp[target]
}
