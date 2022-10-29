package problems

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	var backTracking func(nums []int, path []int)
	backTracking = func(nums, path []int) {
		if len(nums) == 0 {
			result = append(result, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			num := nums[i]
			deletedNums := make([]int, len(nums))
			copy(deletedNums, nums)
			deletedNums = append(deletedNums[:i], deletedNums[i+1:]...)
			backTracking(deletedNums, append(path, num))
		}
	}
	backTracking(nums, []int{})
	return result
}
