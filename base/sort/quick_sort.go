package sort

func QuickSort(a []int) []int {
	partition := func(a []int, lo, hi int) int {
		i, j := lo+1, hi
		for {
			for ; i < hi && a[i] <= a[lo]; i++ {
			}
			for ; j > lo && a[j] >= a[lo]; j-- {
			}
			if j <= i {
				break
			}
			a[i], a[j] = a[j], a[i]
		}
		a[i-1], a[lo] = a[lo], a[i-1]
		return i - 1
	}

	var sort func(a []int, lo, hi int)
	sort = func(a []int, lo, hi int) {
		if lo >= hi {
			return
		}
		j := partition(a, lo, hi)
		sort(a, lo, j-1)
		sort(a, j+1, hi)
	}
	sort(a, 0, len(a)-1)
	return a
}
