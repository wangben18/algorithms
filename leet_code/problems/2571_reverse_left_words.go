package problems

func reverseLeftWords(s string, n int) string {
	ss := []byte(s)
	reverse(ss, 0, n-1)
	reverse(ss, n, len(ss)-1)
	reverse(ss, 0, len(ss)-1)
	return string(ss)
}
