package problems

// 单调栈方式
func trap(height []int) int {
	ans := 0
	stack := []int{0}
	for i := 1; i < len(height); i++ {
		for len(stack) > 0 && height[stack[len(stack)-1]] < height[i] {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				left := stack[len(stack)-1]
				ans += (min(height[left], height[i]) - height[mid]) * (i - left - 1)
			}
		}
		stack = append(stack, i)
	}

	return ans
}

// 双指针方式
// func trap(height []int) int {
// 	ans := 0
// 	rMaxHeight := make([]int, len(height))
// 	rMaxHeight[len(rMaxHeight)-1] = height[len(height)-1]
// 	for i := len(rMaxHeight) - 2; i >= 0; i-- {
// 		rMaxHeight[i] = max(height[i], rMaxHeight[i+1])
// 	}
// 	lMaxHeight := make([]int, len(height))
// 	lMaxHeight[0] = height[0]
// 	for i := 1; i < len(lMaxHeight); i++ {
// 		lMaxHeight[i] = max(height[i], lMaxHeight[i-1])
// 	}

// 	for i, h := range height {
// 		ans += min(lMaxHeight[i], rMaxHeight[i]) - h
// 	}

// 	return ans
// }
