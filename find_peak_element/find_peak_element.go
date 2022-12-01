package find_peak_element

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (right-left)/2 + left
		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
