package max_product

import "fmt"

func _assertTruef(v bool, p string, val ...interface{}) {
	if !v {
		panic(fmt.Errorf(p, val...))
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func maxProduct(nums []int) int {
	_assertTruef(len(nums) != 0, "empty input %+v", nums)
	preMax, preMin := nums[0], nums[0]
	result := preMax
	for i := 1; i < len(nums); i++ {
		curMax := max(preMin*nums[i], max(nums[i], preMax*nums[i]))
		curMin := min(preMin*nums[i], min(nums[i], preMax*nums[i]))
		if curMax > result {
			result = curMax
		}
		preMax, preMin = curMax, curMin
	}
	return result
}
