package problems

func combinationSum3(k int, n int) [][]int {
	result := make([][]int, 0)
	var backTracking func(k, n, start int, path []int)
	backTracking = func(k, n, start int, path []int) {
		if len(path) == k {
			if n == 0 {
				result = append(result, append([]int{}, path...))
			}
			return
		}
		for i := start; i <= 9; i++ {
			backTracking(k, n-i, i+1, append(path, i))
		}
	}
	backTracking(k, n, 1, []int{})
	return result
}
