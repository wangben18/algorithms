package problems

// 动态规划
// 1. dp[i][j] 表示 text1[0:i] 与 text2[0:j] 的最长公共子序列
// 2. 推导方式：分两种情况（text1[i-1]与text2[j-1]相不相同）
//    相同时：说明又多了一个公共字符，dp[i][j] = dp[i-1][j-1] + 1
// 	  不同时：说明不比之前多公共字符，所以比较 dp[i-1][j] 与 dp[i][j-1]，取最大值
// 3. 初始化：为 0 即可
// 4. 推导顺序：从前往后
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}
	for i := 1; i < len(text1)+1; i++ {
		for j := 1; j < len(text2)+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}
