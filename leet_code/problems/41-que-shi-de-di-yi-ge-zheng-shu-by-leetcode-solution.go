package problems

func firstMissingPositive(nums []int) int {
	for i := range nums {
		if nums[i] <= 0 {
			nums[i] = len(nums) + 1
		}
	}
	for _, v := range nums {
		if v < 0 {
			v = -v
		}
		if v <= len(nums) && nums[v-1] > 0 {
			nums[v-1] = -nums[v-1]
		}
	}
	for i, v := range nums {
		if v > 0 {
			return i + 1
		}
	}
	return len(nums) + 1
}
