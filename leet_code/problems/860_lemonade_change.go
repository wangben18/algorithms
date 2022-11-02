package problems

func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for _, num := range bills {
		if num == 5 {
			five++
		} else if num == 10 {
			if five == 0 {
				return false
			}
			five--
			ten++
		} else if num == 20 {
			if ten > 0 && five > 0 {
				ten--
				five--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}
