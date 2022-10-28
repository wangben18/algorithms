package problems

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	var backTracking func(nums []int, start int, path []int)
	backTracking = func(nums []int, start int, path []int) {
		result = append(result, append([]int{}, path...))
		for i := start; i < len(nums); i++ {
			backTracking(nums, i+1, append(path, nums[i]))
		}
	}
	backTracking(nums, 0, []int{})
	return result
}
