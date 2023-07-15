package problems

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))

	stack := []int{0}

	for i := 1; i < len(temperatures); i++ {
		for {
			if len(stack)-1 < 0 || temperatures[stack[len(stack)-1]] >= temperatures[i] {
				break
			}
			result[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return result
}
