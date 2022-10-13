package rob

func getArrayIdx(nums []int, idx, val int) int {
	if idx >= 0 && idx < len(nums) {
		return nums[idx]
	}
	return val
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func rob(nums []int) int {
	var result int
	mark := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		cur := max(
			nums[i]+getArrayIdx(mark, i-2, 0),
			getArrayIdx(mark, i-1, 0),
		)
		if cur > result {
			result = cur
		}
		mark = append(mark, cur)
	}
	return result
}
