package problems

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	anagram := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if num, ok := anagram[s[i]]; ok {
			anagram[s[i]] = num + 1
		} else {
			anagram[s[i]] = 1
		}
	}
	for i := 0; i < len(t); i++ {
		if num, ok := anagram[t[i]]; ok && num > 0 {
			anagram[t[i]] = num - 1
		} else {
			return false
		}
	}
	return true
}
