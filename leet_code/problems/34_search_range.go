package problems

func searchRange(nums []int, target int) []int {
	leftBorder := searchLeftBorder(nums, target)
	rightBorder := searchRightBorder(nums, target)
	if leftBorder == -2 || rightBorder == -2 {
		return []int{-1, -1}
	}
	if rightBorder-leftBorder <= 1 {
		return []int{-1, -1}
	}
	return []int{leftBorder + 1, rightBorder - 1}
}

func searchLeftBorder(nums []int, target int) int {
	var (
		left       = 0
		right      = len(nums) - 1
		leftBorder = -2
	)
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
			leftBorder = right
		}
	}
	return leftBorder
}

func searchRightBorder(nums []int, target int) int {
	var (
		left        = 0
		right       = len(nums) - 1
		rightBorder = -2
	)
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
			rightBorder = left
		}
	}
	return rightBorder
}
