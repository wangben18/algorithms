package problems

// 动态规划
// 1. dp[i][j] 表示 s[:j] 中有多少种 t[:i] 的组成方式
// 2. 推导：分两种情况
//       若 t[i-1] == s[j-1]: 可选择是否使用 s[j-1]，使两种情况相加，即 dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
//       若 t[i-1] != s[j-1]: 用不上 s[j-1]，所以延续 dp[i][j-1]，即dp[i][j] = dp[i][j-1]
// 3. 初始化：空字符串在空字符串中有一种组成方式，所以dp[0][:] 均设为 1
// t 比 s 短，将 t 纵置将 s 横置使内存存取效率更高
func numDistinct(s string, t string) int {
	dp := make([][]int, len(t)+1)
	for i := range dp {
		dp[i] = make([]int, len(s)+1)
	}
	for i := range dp[0] {
		dp[0][i] = 1
	}
	for i := 1; i <= len(t); i++ {
		for j := 1; j <= len(s); j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	return dp[len(t)][len(s)]
}
