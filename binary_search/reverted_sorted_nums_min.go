package binary_search

import "fmt"

func _assertTruef(v bool, f string, fields ...interface{}) {
	if !v {
		panic(fmt.Errorf(f, fields...))
	}
}

func findMin(nums []int) int {
	_assertTruef(len(nums) != 0, "empty input")
	if len(nums) == 1 {
		return nums[0]
	}
	left, right := 0, len(nums)-1
	for left < right {
		if left+1 == right {
			if nums[left] > nums[right] {
				return nums[right]
			}
			return nums[0]
		}
		mid := (right-left)/2 + left
		if nums[left] < nums[mid] {
			left = mid
		} else {
			right = mid
		}
	}
	panic("invalid input")
}
