package problems

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	result, currentMaxDistance, nextMaxDistance := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		if nextDistance := nums[i] + i; nextDistance > nextMaxDistance {
			nextMaxDistance = nums[i] + i
		}
		if i == currentMaxDistance {
			if currentMaxDistance != len(nums)-1 {
				result++
				currentMaxDistance = nextMaxDistance
				if nextMaxDistance >= len(nums)-1 {
					break
				}
			} else {
				break
			}
		}
	}

	return result
}
