package problems

func sortArrayByParityII(nums []int) []int {
	oddIndex := 1
	for i := 0; i < len(nums); i += 2 {
		if nums[i]%2 != 0 {
			for nums[oddIndex]%2 == 1 {
				oddIndex += 2
			}
			nums[i], nums[oddIndex] = nums[oddIndex], nums[i]
		}
	}
	return nums
}
