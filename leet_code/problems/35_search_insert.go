package problems

func SearchInsert(nums []int, target int) int {
	var (
		left  = 0
		right = len(nums)
	)
	for left < right {
		middle := left + ((right - left) >> 1)
		if nums[middle] > target {
			right = middle
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			return middle
		}
	}
	return left
}
