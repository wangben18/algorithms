package problems

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ans := make([]int, len(nums1))
	for i := range ans {
		ans[i] = -1
	}
	numMap := map[int]int{}
	for k, v := range nums1 {
		numMap[v] = k
	}

	stack := []int{nums2[0]}

	for i := 1; i < len(nums2); i++ {
		for {
			topIndex := len(stack) - 1
			if topIndex < 0 || stack[topIndex] >= nums2[i] {
				break
			}
			if num, ok := numMap[stack[topIndex]]; ok {
				ans[num] = nums2[i]
			}
			stack = stack[:topIndex]
		}
		stack = append(stack, nums2[i])
	}

	return ans
}
