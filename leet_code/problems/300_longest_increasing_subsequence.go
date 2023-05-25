package problems

// 动态规划
// 1. dp[i] 定义：表示 i 位置的最长子系列长度
// 2. 初始值：1
// 3. 推导算法：取小于 i 位置数字的最大长度 + 1
// 4. 推导顺序：从头到尾
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	res := 1
	for i := 1; i < len(dp); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				if dp[i] > res {
					res = dp[i]
				}
			}
		}
	}
	return res
}
