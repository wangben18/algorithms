package problems

func combine(n int, k int) [][]int {
	result := make([][]int, 0)
	var backTracking func(n, k, start int, path []int)
	backTracking = func(n, k, start int, path []int) {
		if len(path) == k {
			newPath := make([]int, len(path))
			copy(newPath, path)
			result = append(result, newPath)
			return
		}
		for i := start; i <= n; i++ {
			backTracking(n, k, i+1, append(path, i))
		}
	}
	backTracking(n, k, 1, []int{})

	return result
}
