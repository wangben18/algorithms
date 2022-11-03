package problems

import "strconv"

func monotoneIncreasingDigits(n int) int {
	str := []byte(strconv.Itoa(n))
	nineFlag := len(str)
	for i := len(str) - 1; i >= 1; i-- {
		if str[i-1] > str[i] {
			nineFlag = i
			str[i-1]--
		}
	}
	for i := len(str) - 1; i >= nineFlag; i-- {
		str[i] = '9'
	}
	result, _ := strconv.Atoi(string(str))
	return result
}
