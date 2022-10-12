package problems

type HumdrumQueue struct {
	slice []int
}

func (this *HumdrumQueue) Pop(x int) {
	if len(this.slice) > 0 && this.Front() == x {
		this.slice = this.slice[1:]
	}
}

func (this *HumdrumQueue) Push(x int) {
	for len(this.slice) > 0 && x > this.slice[len(this.slice)-1] {
		this.slice = this.slice[:len(this.slice)-1]
	}
	this.slice = append(this.slice, x)
}

func (this *HumdrumQueue) Front() int {
	return this.slice[0]
}

func maxSlidingWindow(nums []int, k int) []int {
	humdrumQueue := HumdrumQueue{}
	max := make([]int, 0, len(nums))
	for i := 0; i < k; i++ {
		humdrumQueue.Push(nums[i])
	}
	max = append(max, humdrumQueue.Front())
	for i := k; i < len(nums); i++ {
		humdrumQueue.Pop(nums[i-k])
		humdrumQueue.Push(nums[i])
		max = append(max, humdrumQueue.Front())
	}
	return max
}
