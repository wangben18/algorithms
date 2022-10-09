package problems

func canConstruct(ransomNote string, magazine string) bool {
	counts := make([]int, 26)
	for _, m := range magazine {
		counts[m-'a']++
	}
	for _, r := range ransomNote {
		counts[r-'a']--
		if counts[r-'a'] < 0 {
			return false
		}
	}
	return true
}
