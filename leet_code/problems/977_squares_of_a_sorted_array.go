package problems

func sortedSquares(nums []int) []int {
	left := 0
	right := len(nums) - 1
	sorted := make([]int, len(nums))
	for sortedIndex := len(nums) - 1; sortedIndex >= 0; sortedIndex-- {
		leftSquare := nums[left] * nums[left]
		rightSquare := nums[right] * nums[right]
		if leftSquare < rightSquare {
			sorted[sortedIndex] = rightSquare
			right--
		} else {
			sorted[sortedIndex] = leftSquare
			left++
		}
	}
	return sorted
}
