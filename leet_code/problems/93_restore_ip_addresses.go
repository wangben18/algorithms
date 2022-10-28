package problems

import (
	"strconv"
)

func restoreIpAddresses(s string) []string {
	if s == "" {
		return []string{}
	}
	result := make([]string, 0)
	isValidAddressNum := func(s string) (bool, bool) {
		if len(s) > 1 && s[0] == '0' {
			return false, false
		}
		if num, err := strconv.ParseInt(s, 10, 64); err != nil {
			return false, true
		} else if num > 255 {
			return false, false
		}
		return true, false
	}
	var backTracking func(s string, start, dotCount int)
	backTracking = func(s string, start, dotCount int) {
		if dotCount == 3 {
			if isValid, _ := isValidAddressNum(s[start:]); isValid {
				result = append(result, s)
			}
			return
		}
		for i := start; i < len(s) && len(s[i:])/3.0 <= 3-dotCount; i++ {
			numStr := s[start:i]
			if isValid, isContinue := isValidAddressNum(numStr); !isValid && !isContinue {
				break
			} else if !isValid && isContinue {
				continue
			}
			backTracking(s[:i]+"."+s[i:], i+1, dotCount+1)
		}
	}
	backTracking(s, 0, 0)
	return result
}
