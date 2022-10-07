package problems

func intersection(nums1 []int, nums2 []int) []int {
	intersectionNums := make([]int, 0)
	intersectionMap := make(map[int]struct{})
	for _, num := range nums1 {
		if _, ok := intersectionMap[num]; !ok {
			intersectionMap[num] = struct{}{}
		}
	}
	for _, num := range nums2 {
		if _, ok := intersectionMap[num]; ok {
			delete(intersectionMap, num)
			intersectionNums = append(intersectionNums, num)
		}
	}
	return intersectionNums
}
