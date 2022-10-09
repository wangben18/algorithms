package problems

func reverseWords(s string) string {
	ss := []byte(s)
	ss = removeExtraSpace(ss)
	reverse(ss, 0, len(ss)-1)
	start := 0
	for i := 0; i <= len(ss); i++ {
		if i == len(ss) || ss[i] == ' ' {
			reverse(ss, start, i-1)
			start = i + 1
		}
	}
	return string(ss)
}

func removeExtraSpace(ss []byte) []byte {
	slow := 0
	for fast := 0; fast < len(ss); fast++ {
		if ss[fast] != ' ' {
			if slow > 0 {
				ss[slow] = ' '
				slow++
			}
			for fast < len(ss) && ss[fast] != ' ' {
				ss[slow] = ss[fast]
				slow++
				fast++
			}
		}
	}
	return ss[0:slow]
}

// 反转字符串（左闭右闭）
func reverse(ss []byte, begin, end int) {
	for begin < end {
		ss[begin], ss[end] = ss[end], ss[begin]
		begin++
		end--
	}
	return
}
