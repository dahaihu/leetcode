package gas_station

func gas(gas []int, cost []int) int {
	// (start, end]
	start, end := len(cost)-1, 0
	cur := gas[end] - cost[end]
	for start > end {
		if cur >= 0 {
			end += 1
			cur += gas[end] - cost[end]
		} else {
			cur += gas[start] - cost[start]
			start -= 1
		}
	}
	if cur >= 0 {
		return (start + 1) % len(cost)
	}
	return -1
}
