package three_sum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		start, end := i+1, len(nums)-1
		for start < end {
			cur := nums[i] + nums[start] + nums[end]
			switch {
			case cur == 0:
				result = append(result, []int{nums[i], nums[start], nums[end]})
				start += 1
				for start < end && nums[start] == nums[start-1] {
					start++
				}
				end -= 1
				for start < end && nums[end] == nums[end+1] {
					end -= 1
				}
			case cur > 0:
				end -= 1
			case cur < 0:
				start += 1
			}
		}
	}
	return result
}
