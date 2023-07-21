package problems

func isLongPressedName(name string, typed string) bool {
	if name[0] != typed[0] {
		return false
	}
	nameIdx := 1
	for i := 1; i < len(typed); i++ {
		if nameIdx == len(name) {
			if typed[i] != typed[i-1] {
				return false
			}
		} else {
			if name[nameIdx] != typed[i] && typed[i] != typed[i-1] {
				return false
			} else if name[nameIdx] == typed[i] {
				nameIdx++
			}
		}
	}

	return nameIdx == len(name)
}
