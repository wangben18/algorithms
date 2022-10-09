package problems

func replaceSpace(s string) string {
	ss := []byte(s)
	length := len(ss)
	spaceCount := 0
	for _, b := range ss {
		if b == ' ' {
			spaceCount++
		}
	}
	ss = append(ss, make([]byte, spaceCount*2)...)
	newLength := len(ss)
	l := length - 1
	r := newLength - 1
	for l < r {
		if ss[l] == ' ' {
			ss[r] = '0'
			ss[r-1] = '2'
			ss[r-2] = '%'
			l--
			r -= 3
		} else {
			ss[r] = ss[l]
			r--
			l--
		}
	}
	return string(ss)
}
