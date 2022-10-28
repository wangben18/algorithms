package problems

func partition(s string) [][]string {
	result := make([][]string, 0)
	isPalindrome := func(s string) bool {
		for l, r := 0, len(s)-1; l < r; l++ {
			if s[l] != s[r] {
				return false
			}
			r--
		}
		return true
	}
	var backTracking func(s string, start int, path []string)
	backTracking = func(s string, start int, path []string) {
		if start == len(s) {
			result = append(result, append([]string{}, path...))
			return
		}
		for i := start; i < len(s); i++ {
			if !isPalindrome(s[start : i+1]) {
				continue
			}
			backTracking(s, i+1, append(path, s[start:i+1]))
		}
	}
	backTracking(s, 0, []string{})
	return result
}
