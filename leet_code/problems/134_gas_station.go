package problems

func canCompleteCircuit(gas []int, cost []int) int {
	min := gas[0] - cost[0]
	sum := 0
	for i := 0; i < len(gas); i++ {
		remain := gas[i] - cost[i]
		sum += remain
		if remain < min {
			remain = min
		}
	}
	if min >= 0 {
		return 0
	}
	if sum < 0 {
		return -1
	}
	for i := len(gas) - 1; i >= 0; i-- {
		remain := gas[i] - cost[i]
		min += remain
		if min >= 0 {
			return i
		}
	}
	return -1
}
