package rob

import "fmt"

func _assertTruef(v bool, f string, fields ...interface{}) {
	if !v {
		panic(fmt.Errorf(f, fields...))
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func rob(nums []int) int {
	_assertTruef(len(nums) != 0, "empty input")
	pre, prepre := 0, 0
	var out int
	for _, num := range nums {
		cur := max(num+prepre, pre)
		if cur > out {
			out = cur
		}
		pre, prepre = cur, pre
	}
	return out
}
