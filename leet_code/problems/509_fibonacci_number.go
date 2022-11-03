package problems

func fib(n int) int {
	if n <= 1 {
		return n
	}
	dp := [2]int{0, 1}
	for i := 2; i <= n; i++ {
		dp[0], dp[1] = dp[1], dp[0]+dp[1]
	}
	return dp[1]
}
