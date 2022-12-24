package binary_search

func search(nums []int, target int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	if n == 1 {
		return nums[0] == target
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			return true
		}
		if nums[mid] > nums[left] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] < nums[right] {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if nums[mid] == nums[left] {
				left++
			} else {
				right--
			}
		}
	}

	return false
}

func binaryConvertedMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		if left+1 == right {
			if nums[left] > nums[right] {
				return nums[right]
			}
			return nums[0]
		}
		mid := (right-left)/2 + left
		if nums[left] <= nums[mid] {
			left = mid
		} else {
			right = mid
		}
	}
	return nums[0]
}
