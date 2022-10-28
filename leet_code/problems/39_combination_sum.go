package problems

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := make([][]int, 0)
	var backTracking func(candidates []int, target, start, sum int, path []int)
	backTracking = func(candidates []int, targe, start, sum int, path []int) {
		if sum == target {
			result = append(result, append([]int{}, path...))
			return
		} else if sum > target {
			return
		}
		for i := start; (i < len(candidates)) && (sum+candidates[i] <= target); i++ {
			backTracking(candidates, target, i, sum+candidates[i], append(path, candidates[i]))
		}
	}
	backTracking(candidates, target, 0, 0, []int{})
	return result
}
