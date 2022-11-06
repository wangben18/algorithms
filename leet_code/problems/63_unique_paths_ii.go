package problems

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	rowCount := len(obstacleGrid)
	columnCount := len(obstacleGrid[0])
	dp := make([][]int, rowCount)
	for i := range dp {
		dp[i] = make([]int, columnCount)
	}
	for row := 0; row < rowCount; row++ {
		for col := 0; col < columnCount; col++ {
			if obstacleGrid[row][col] == 0 {
				if row == 0 && col == 0 {
					dp[row][col] = 1
				} else if row == 0 {
					if dp[row][col-1] == 0 {
						dp[row][col] = 0
					} else {
						dp[row][col] = 1
					}
				} else if col == 0 {
					if dp[row-1][col] == 0 {
						dp[row][col] = 0
					} else {
						dp[row][col] = 1
					}
				} else {
					dp[row][col] = dp[row-1][col] + dp[row][col-1]
				}
			}
		}
	}
	return dp[rowCount-1][columnCount-1]
}
