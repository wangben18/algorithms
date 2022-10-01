package problems

func minSubArrayLen(target int, nums []int) int {
	left := 0
	sum := 0
	minLen := 0
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum >= target {
			if right-left+1 < minLen || minLen == 0 {
				minLen = right - left + 1
			}
			sum -= nums[left]
			left++
		}
	}
	return minLen
}
