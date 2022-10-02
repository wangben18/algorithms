package problems

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	num := 1
	for padding := 0; padding <= n/2; padding++ {
		for i := padding; i < n-1-padding; i++ {
			matrix[padding][i] = num
			num++
		}
		for i := padding; i < n-1-padding; i++ {
			matrix[i][n-1-padding] = num
			num++
		}
		for i := padding; i < n-1-padding; i++ {
			matrix[n-1-padding][n-1-i] = num
			num++
		}
		for i := padding; i < n-1-padding; i++ {
			matrix[n-1-i][padding] = num
			num++
		}
	}
	if n%2 == 1 {
		matrix[n/2][n/2] = num
	}

	return matrix
}
