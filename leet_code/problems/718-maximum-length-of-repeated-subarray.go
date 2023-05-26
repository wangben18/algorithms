package problems

// 动态规划
// 1. dp[i][j] 表示 nums1[0:i-1] 与 nums2[0:j-1] 并以最后字符为结尾的最长重复子数组长度
// 2. dp 初始化为 0
// 3. 推导方式：nums1[i-1] 与 nums2[j-1]相等时，dp[i][j] 比dp[i-1][j-1] 加 1 长度
func findLength(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}
	result := 0
	for i := 1; i < len(nums1)+1; i++ {
		for j := 1; j < len(nums2)+1; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > result {
					result = dp[i][j]
				}
			}
		}
	}
	return result
}
