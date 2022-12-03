package sorted

func reverseRange(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func sortedPow(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left] >= 0 {
			break
		}
		if nums[right] < 0 {
			reverseRange(nums, left, right)
			break
		}
		if abs(nums[left]) < nums[right] {
			right--
		} else {
			oldLeft := nums[left]
			copy(nums[left:right], nums[left+1:right+1])
			nums[right] = oldLeft
			right--
		}
	}
	return nums
}
