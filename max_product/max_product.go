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
	preMax, preMin, out := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		cur := nums[i]
		curMax := max(cur, max(preMax*cur, preMin*cur))
		curMin := min(cur, min(preMax*cur, preMin*cur))
		if curMax > out {
			out = curMax
		}
		preMax, preMin = curMax, curMin
	}
	return out
}
