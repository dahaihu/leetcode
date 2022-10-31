package quick_sort

func position(nums []int, start, end int) int {
	target := nums[start]
	left, right := start+1, end
	for left <= right {
		for left <= right && nums[left] <= target {
			left += 1
		}
		for left <= right && nums[right] >= target {
			right -= 1
		}
		if left > right {
			break
		}
		nums[left], nums[right] = nums[right], nums[left]
		left += 1
		right -= 1
	}
	nums[start], nums[right] = nums[right], nums[start]
	return right
}

func quickSort(nums []int, start, end int) {
	if start < end {
		mid := position(nums, start, end)
		quickSort(nums, start, mid-1)
		quickSort(nums, mid+1, end)
	}
}
