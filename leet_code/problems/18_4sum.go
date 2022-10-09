package problems

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums)-3; i++ {
		ni := nums[i]
		if i > 0 && nums[i-1] == ni {
			continue
		}
		if nums[i] > target && target > 0 {
			break
		}
		for j := i + 1; j < len(nums)-2; j++ {
			nj := nums[j]
			if j > i+1 && nums[j-1] == nj {
				continue
			}
			if nums[i]+nums[j] > target && target > 0 {
				break
			}
			l, r := j+1, len(nums)-1
			for l < r {
				nl, nr := nums[l], nums[r]
				sum := ni + nj + nl + nr
				if sum == 0 {
					result = append(result, []int{ni, nj, nl, nr})
					for l < r && nums[l] == nl {
						l++
					}
					for l < r && nums[r] == nr {
						r--
					}
				} else if sum < 0 {
					l++
				} else {
					r--
				}
			}
		}
	}
	return result
}
