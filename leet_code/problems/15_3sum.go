package problems

import "sort"

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for r > l {
			n, ln, rn := nums[i], nums[l], nums[r]
			sum := n + ln + rn
			if sum == 0 {
				result = append(result, []int{n, ln, rn})
				for r > l && nums[l] == ln {
					l++
				}
				for r > l && nums[r] == rn {
					r--
				}
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return result
}
