package problems

func partitionLabels(s string) []int {
	charMap := make([]int, 26)
	for i, c := range s {
		charMap[byte(c)-'a'] = i
	}
	max := -1
	result := make([]int, 0)
	pre := -1
	for i, c := range s {
		if charMap[byte(c)-'a'] > max {
			max = charMap[byte(c)-'a']
		}
		if i == max {
			result = append(result, i-pre)
			pre = max
		}
	}
	return result
}
