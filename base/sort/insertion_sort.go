package sort

func InsertionSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	for i := 1; i < len(a); i++ {
		for j := i; j > 0 && a[j] < a[j-1]; j-- {
			a[j-1], a[j] = a[j], a[j-1]
		}
	}
	return a
}
