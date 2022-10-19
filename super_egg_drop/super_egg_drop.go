package super_egg_drop

import "math"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// superEggDrop k egg, n floor
func superEggDrop(k int, n int) [][]int {
	mark := make([][]int, k)
	for i := 0; i < k; i++ {
		mark[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		mark[0][i] = i + 1
	}
	for i := 0; i < k; i++ {
		mark[i][0] = 1
	}
	for egg := 1; egg < k; egg++ {
		for floor := 1; floor < n; floor++ {
			val := math.MaxInt64
			for innerFloor := 0; innerFloor <= floor; innerFloor++ {
				var breakCondition int
				if innerFloor >= 1 {
					breakCondition = mark[egg-1][innerFloor-1]
				}
				var noBreakCondition int
				if innerFloor < floor {
					noBreakCondition = mark[egg][floor-innerFloor-1]
				}
				if cur := max(breakCondition, noBreakCondition); cur < val {
					val = cur
				}
			}
			mark[egg][floor] = val + 1
		}
	}
	return mark
}
