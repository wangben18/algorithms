package problems

func letterCasePermutation(s string) []string {
	result := make([]string, 0)
	var backTracking func(s string, start int, path string)
	backTracking = func(s string, start int, path string) {
		if len(path) == len(s) {
			result = append(result, path)
			return
		}
		for i := start; i < len(s) && len(path) == i; i++ {
			if s[i] >= '0' && s[i] <= '9' {
				path += string(s[i])
				if len(path) == len(s) {
					result = append(result, path)
					return
				}
				continue
			}
			offset := 'a' - 'A'
			if s[i] >= 'a' && s[i] <= 'z' {
				offset = 'A' - 'a'
			}
			backTracking(s, i+1, path+string(s[i]))
			backTracking(s, i+1, path+string(s[i]+byte(offset)))
		}
	}
	backTracking(s, 0, "")
	return result
}
