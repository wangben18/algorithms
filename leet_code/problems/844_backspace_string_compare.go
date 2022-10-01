package problems

func backspaceCompare(s string, t string) bool {
	return backspaceString(s) == backspaceString(t)
}

func backspaceString(s string) string {
	byteS := []byte(s)
	slow := 0
	for fast := 0; fast < len(s); fast++ {
		if byteS[fast] == '#' {
			if slow > 0 {
				slow--
			}
			continue
		}
		byteS[slow] = byteS[fast]
		slow++
	}
	return string(byteS[:slow])
}
