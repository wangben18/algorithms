package problems

func binarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	var (
		left  = 0
		right = len(nums)
	)
	for left < right {
		middle := (left + right) / 2
		if nums[middle] > target {
			right = middle
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			return middle
		}
	}
	return -1
}
