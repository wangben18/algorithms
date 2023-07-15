package problems

func nextGreaterElements(nums []int) []int {
	ans := make([]int, len(nums))
	for i := range ans {
		ans[i] = -1
	}
	stack := []int{0}

	for i := 1; i < len(nums)*2; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%len(nums)] {
			ans[stack[len(stack)-1]] = nums[i%len(nums)]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%len(nums))
	}

	return ans
}
