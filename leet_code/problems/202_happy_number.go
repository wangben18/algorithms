package problems

func isHappy(n int) bool {
	numMap := make(map[int]bool)
	for n != 1 && !numMap[n] {
		n, numMap[n] = getSum(n), true
	}
	return n == 1
}

func getSum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}
