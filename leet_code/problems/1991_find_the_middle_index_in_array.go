package problems

func PivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return 0
	}
	total := 0
	for _, num := range nums {
		total += num
	}
	currentTotal := 0
	for i, num := range nums {
		currentTotal += num
		if total-currentTotal == currentTotal-num {
			return i
		}
	}
	return -1
}
