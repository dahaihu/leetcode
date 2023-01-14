package max_product

import "fmt"

func _assertTruef(v bool, p string, val ...interface{}) {
	if !v {
		panic(fmt.Errorf(p, val...))
	}
}

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

func maxProduct(nums []int) int {
	_assertTruef(len(nums) != 0, "empty input")
	out := nums[0]
	preMin, preMax := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		cur := nums[i]
		nextPreMax := max(cur, max(preMin*cur, preMax*cur))
		nextPreMin := min(cur, min(preMin*cur, preMax*cur))
		preMin, preMax = nextPreMin, nextPreMax
		if preMax > out {
			out = preMax
		}
	}
	return out
}
