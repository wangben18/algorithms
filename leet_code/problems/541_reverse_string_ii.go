package problems

func reverseStr(s string, k int) string {
	ss := []byte(s)
	length := len(ss)
	for i := 0; i < length; i += 2 * k {
		if i+k < length {
			reverseString(ss[i : i+k])
		} else {
			reverseString(ss[i:length])
		}
	}
	return string(ss)
}
