package problems

func commonChars(words []string) []string {
	hash := make([][]int, len(words))
	for i, w := range words {
		hash[i] = make([]int, 26)
		for j := 0; j < len(w); j++ {
			hash[i][w[j]-'a']++
		}
	}
	result := hash[0]

	for i := 1; i < len(hash); i++ {
		for j, v := range hash[i] {
			if v < result[j] {
				result[j] = v
			}
		}
	}
	ans := make([]string, 0)
	for i, v := range result {
		if v > 0 {
			for j := 0; j < v; j++ {
				ans = append(ans, string(append([]byte{}, byte('a'+i))))
			}
		}
	}
	return ans
}
