package find_peak_element

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)
	for left < right {
		if left+1 == right {
			return left
		}
		mid := (right-left)/2 + left
		if nums[mid-1] <= nums[mid] {
			left = mid
		} else {
			right = mid
		}
	}
	panic("invalid input")
}
