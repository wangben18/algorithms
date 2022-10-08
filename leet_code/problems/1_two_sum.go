package problems

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		numMap[num] = i
	}
	for i, num := range nums {
		if j, ok := numMap[target-num]; ok && i != j {
			return []int{i, j}
		}
	}
	return nil
}
