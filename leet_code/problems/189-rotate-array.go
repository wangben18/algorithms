package problems

func rotate(nums []int, k int) {
	reverse := func(n []int) {
		for i := 0; i <= len(n)/2; i++ {
			n[i], n[len(n)-i-1] = n[len(n)-i-1], n[i]
		}
	}
	reverse(nums)
	reverse(nums[:k%len(nums)])
	reverse(nums[k%len(nums):])
}
