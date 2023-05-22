package sort

func MergeSort(a []int) []int {
	helper := make([]int, len(a))
	merge := func(a []int, lo, mid, hi int) {
		for k := lo; k <= hi; k++ {
			helper[k] = a[k]
		}
		i, j := lo, mid+1
		for k := lo; k <= hi; k++ {
			if i > mid {
				a[k] = helper[j]
				j++
			} else if j > hi {
				a[k] = helper[i]
				i++
			} else if helper[i] <= helper[j] {
				a[k] = helper[i]
				i++
			} else {
				a[k] = helper[j]
				j++
			}
		}
	}
	var sort func(a []int, lo, hi int)
	sort = func(a []int, lo, hi int) {
		if lo >= hi {
			return
		}
		mid := lo + (hi-lo)/2
		sort(a, lo, mid)
		sort(a, mid+1, hi)
		merge(a, lo, mid, hi)
	}
	sort(a, 0, len(a)-1)
	return a
}
