package problems

// 1. dp[i][j] 表示 word1[:i] 与 word2[:j] 之间最少编辑距离
// 2. 推导：两种情况
//     若 word1[i-1] == word2[j-1]：则最后一个字段相等，无需任何操作，即 dp[i][j] = dp[i-1][j-1]
//     若不相等：有三种操作
//           删：可以是 word1 删去或 word2 删去，即 dp[i-1][j] + 1 或 dp[i][j-1] + 1
//           增：删 word1 也可以改为增 word2，所以公式与删相同
//           替换：将其中一个 word 最后一个字段替换为另一个字段，即 dp[i-1][j-1] + 1
func editDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 1; i <= len(word1); i++ {
		dp[i][0] = i
	}
	for i := 1; i <= len(word2); i++ {
		dp[0][i] = i
	}
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j])) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}
