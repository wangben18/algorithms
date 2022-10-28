package problems

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	var backTracking func(nums []int, start int, path []int)
	backTracking = func(nums []int, start int, path []int) {
		result = append(result, append([]int{}, path...))
		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			backTracking(nums, i+1, append(path, nums[i]))
		}
	}
	backTracking(nums, 0, []int{})
	return result
}
