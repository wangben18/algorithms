package problems

func productExceptSelf(nums []int) []int {
	left := make([]int, len(nums))
	left[0] = 1
	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	right := 1
	for i := len(nums) - 1; i >= 0; i-- {
		left[i] *= right
		right *= nums[i]
	}
	return left
}
