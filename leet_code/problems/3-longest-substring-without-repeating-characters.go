package problems

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	longest := 0
	left, right := 0, 0
	m := make(map[byte]struct{})

	length := 0
	for left < len(s) && right < len(s) {
		if _, ok := m[s[right]]; ok {
			delete(m, s[left])
			left++
			length--
		} else {
			m[s[right]] = struct{}{}
			right++
			length++
			if length > longest {
				longest = length
			}
		}
	}

	return longest
}
