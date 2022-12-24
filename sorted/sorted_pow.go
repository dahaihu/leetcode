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

func sorted2powSorted(nums []int) []int {
	right := len(nums) - 1
	for right > 0 {
		if nums[0] >= 0 {
			break
		}
		if nums[right] < 0 {
			reverseRange(nums, 0, right)
			break
		}
		if abs(nums[0]) < nums[right] {
			right--
		} else {
			oldLeft := nums[0]
			copy(nums[:right], nums[1:right+1])
			nums[right] = oldLeft
			right--
		}
	}
	return nums
}
