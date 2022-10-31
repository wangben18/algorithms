package problems

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	result, index := 0, 0
	for i := 0; i < len(s); i++ {
		if index < len(g) && g[index] <= s[i] {
			result++
			index++
		}
	}
	return result
}
