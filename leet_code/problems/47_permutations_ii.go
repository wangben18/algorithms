package problems

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	var backTracking func(nums, path []int)
	backTracking = func(nums, path []int) {
		if len(nums) == 0 {
			result = append(result, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			deletedNums := make([]int, len(nums))
			copy(deletedNums, nums)
			deletedNums = append(deletedNums[:i], deletedNums[i+1:]...)
			backTracking(deletedNums, append(path, nums[i]))
		}
	}
	backTracking(nums, []int{})
	return result
}
