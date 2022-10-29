package problems

func findSubsequences(nums []int) [][]int {
	result := make([][]int, 0)
	var backTracking func(nums []int, start int, path []int)
	backTracking = func(nums []int, start int, path []int) {
		if len(path) > 1 {
			result = append(result, append([]int{}, path...))
		}
		usedNums := make([]int, 0)
		for i := start; i < len(nums); i++ {
			if len(path) > 0 && path[len(path)-1] > nums[i] {
				continue
			}
			used := false
			for _, num := range usedNums {
				if num == nums[i] {
					used = true
				}
			}
			if used {
				continue
			}
			usedNums = append(usedNums, nums[i])
			backTracking(nums, i+1, append(path, nums[i]))
		}
	}
	backTracking(nums, 0, []int{})
	return result
}
