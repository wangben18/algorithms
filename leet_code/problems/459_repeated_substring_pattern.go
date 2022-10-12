package problems

func repeatedSubstringPattern(s string) bool {
	if len(s) <= 1 {
		return false
	}
	next := make([]int, len(s))
	next[0] = 0
	j := 0
	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}

	return next[len(s)-1] != 0 && len(s)%(len(s)-next[len(s)-1]) == 0
}
