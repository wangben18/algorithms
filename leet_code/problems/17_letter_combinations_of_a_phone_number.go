package problems

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	digitMap := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	result := make([]string, 0)
	var backTracking func(digits string, index int, path []byte)
	backTracking = func(digits string, index int, path []byte) {
		if len(path) == len(digits) {
			result = append(result, string(path))
			return
		}
		digit := digits[index] - '0'
		letters := digitMap[digit]
		for i := 0; i < len(letters); i++ {
			backTracking(digits, index+1, append(path, letters[i]))
		}
	}
	backTracking(digits, 0, []byte{})
	return result
}
