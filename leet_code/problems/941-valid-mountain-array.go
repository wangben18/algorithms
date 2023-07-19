package problems

func validMountainArray(arr []int) bool {
	if len(arr) < 3 || arr[0] >= arr[1] {
		return false
	}

	decreased := false
	for i := 2; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			return false
		}
		if arr[i] > arr[i-1] && decreased {
			return false
		}
		if arr[i] < arr[i-1] && !decreased {
			decreased = true
		}
	}
	return decreased
}
