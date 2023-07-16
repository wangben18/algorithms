package problems

// 单调栈
func largestRectangleArea(heights []int) int {
	ans := 0
	heights = append([]int{0}, append(heights, 0)...)
	stack := []int{0}
	for i := 1; i < len(heights); i++ {
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				left := stack[len(stack)-1]
				right := i
				ans = max(ans, (right-left-1)*heights[mid])
			}
		}
		stack = append(stack, i)
	}
	return ans
}

// 双指针
// func largestRectangleArea(heights []int) int {
// 	leftNearLower := make([]int, len(heights))
// 	leftNearLower[0] = -1
// 	for i := 1; i < len(heights); i++ {
// 		t := i - 1
// 		for t >= 0 && heights[t] >= heights[i] {
// 			t = leftNearLower[t]
// 		}
// 		leftNearLower[i] = t
// 	}
// 	rightNearLower := make([]int, len(heights))
// 	rightNearLower[len(rightNearLower)-1] = -1
// 	for i := len(heights) - 1; i >= 0; i-- {
// 		t := i + 1
// 		for t < len(heights) && heights[t] >= heights[i] {
// 			t = rightNearLower[t]
// 		}
// 		rightNearLower[i] = t
// 	}

// 	ans := 0
// 	for i := 0; i < len(heights); i++ {
// 		ans = max(ans, (rightNearLower[i]-leftNearLower[i]-1)*heights[i])
// 	}

// 	return ans
// }
