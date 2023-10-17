package problems

func longestConsecutive(nums []int) int {
	m := make(map[int]struct{})
	for _, v := range nums {
		m[v] = struct{}{}
	}
	longest := 0
	for num := range m {
		if _, ok := m[num-1]; ok {
			continue
		}
		i := 1
		for {
			if _, ok := m[num+i]; !ok {
				break
			}
			i++
		}
		if i > longest {
			longest = i
		}
	}
	return longest
}
