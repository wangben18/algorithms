package problems

import "strconv"

func evalRPN(tokens []string) int {
	calMap := map[string]func(a, b int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}
	stack := make([]string, 0)
	for _, token := range tokens {
		if cal, ok := calMap[token]; ok {
			b, _ := strconv.Atoi(stack[len(stack)-1])
			a, _ := strconv.Atoi(stack[len(stack)-2])
			stack = stack[:len(stack)-2]
			stack = append(stack, strconv.Itoa(cal(a, b)))
			continue
		}
		stack = append(stack, token)
	}
	result, _ := strconv.Atoi(stack[0])
	return result
}
