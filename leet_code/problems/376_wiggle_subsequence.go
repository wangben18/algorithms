package problems

func wiggleMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	result, preDiff := 0, 0
	for i := 0; i < len(nums)-1; i++ {
		curDiff := nums[i+1] - nums[i]
		if (curDiff > 0 && preDiff <= 0) || (curDiff < 0 && preDiff >= 0) {
			preDiff = curDiff
			result++
		}
	}
	return result + 1
}
