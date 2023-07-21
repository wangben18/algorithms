package problems

func isIsomorphic(s string, t string) bool {
	m1, m2 := make(map[byte]byte), make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if v, ok := m1[s[i]]; ok && v != t[i] {
			return false
		}
		if v, ok := m2[t[i]]; ok && v != s[i] {
			return false
		}

		m1[s[i]], m2[t[i]] = t[i], s[i]
	}
	return true
}
