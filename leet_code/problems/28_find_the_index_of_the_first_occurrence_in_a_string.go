package problems

func strStr(haystack string, needle string) int {
	if len(needle) == 0 || len(haystack) == 0 {
		return -1
	}
	j := 0
	next := getNext(needle)
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - len(needle) + 1
		}
	}
	return -1
}

func getNext(s string) []int {
	j := 0
	next := make([]int, len(s))
	next[0] = 0
	for i := 1; i < len(s); i++ {
		for j > 0 && s[j] != s[i] {
			j = next[j-1]
		}
		if s[j] == s[i] {
			j++
		}
		next[i] = j
	}
	return next
}
