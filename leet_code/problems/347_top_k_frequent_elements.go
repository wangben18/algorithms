package problems

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	kMap := make(map[int]int)
	for _, num := range nums {
		kMap[num]++
	}
	h := &IntHeap{}
	heap.Init(h)
	for num, numK := range kMap {
		heap.Push(h, [2]int{num, numK})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	result := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		result[i] = heap.Pop(h).([2]int)[0]
	}
	return result
}

type IntHeap [][2]int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Pop() interface{} {
	val := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return val
}

func (h *IntHeap) Push(val interface{}) {
	*h = append(*h, val.([2]int))
}
