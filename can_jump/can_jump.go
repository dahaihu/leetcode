package can_jump

import "fmt"

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	margin := 0
	for idx, num := range nums {
		if idx > margin {
			return false
		}
		if next := idx + num; next > margin {
			margin = next
		}
		if margin >= len(nums)-1 {
			return true
		}
	}
	return true
}

func canJump2(nums []int) int {
	end, next := 0, 0
	var count int
	for idx, num := range nums {
		if idx > end {
			panic(fmt.Errorf("no reachable idx[%d] in nums[%+v]", idx, nums))
		}
		if idx == len(nums)-1 {
			return count
		}
		if curNext := idx + num; curNext > next {
			next = curNext
		}
		if idx == end {
			end = next
			count++
		}
	}
	panic("not reachable")
}
