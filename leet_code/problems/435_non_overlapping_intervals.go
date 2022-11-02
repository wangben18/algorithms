package problems

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	result, preRightSide := 1, intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if preRightSide <= intervals[i][0] {
			preRightSide = intervals[i][1]
			result++
		}
	}
	return len(intervals) - result
}
