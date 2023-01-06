package problems

// 完全背包
func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= 2; j++ {
			if i-j >= 0 {
				dp[i] += dp[i-j]
			}
		}
	}
	return dp[n]
}

// DP
// func climbStairs(n int) int {
// 	if n <= 2 {
// 		return n
// 	}
// 	dp := make([]int, n)
// 	dp[0] = 1
// 	dp[1] = 2
// 	for i := 2; i < n; i++ {
// 		dp[i] = dp[i-1] + dp[i-2]
// 	}
// 	return dp[n-1]
// }

// 递归（带 memo）
// func climbStairs(n int) int {
// 	return climbStairsWithMemo(n, make(map[int]int))
// }

// func climbStairsWithMemo(n int, memo map[int]int) int {
// 	if n == 1 {
// 		return 1
// 	} else if n == 2 {
// 		return 2
// 	}
// 	if memoNum, ok := memo[n]; ok {
// 		return memoNum
// 	}
// 	num := climbStairsWithMemo(n-1, memo) + climbStairsWithMemo(n-2, memo)
// 	memo[n] = num
// 	return num
// }
