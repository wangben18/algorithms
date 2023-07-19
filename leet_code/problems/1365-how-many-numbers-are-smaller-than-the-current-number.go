package problems

import "sort"

func smallerNumbersThanCurrent(nums []int) []int {
	ans := make([]int, len(nums))
	copy(ans, nums)
	sort.Ints(nums)
	m := map[int]int{}
	for i, v := range nums {
		if _, ok := m[v]; !ok {
			m[v] = i
		}
	}
	for i, v := range ans {
		ans[i] = m[v]
	}

	return ans
}
