package can_jump

import "fmt"

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	margin := nums[0]
	for i := 1; i < len(nums); i++ {
		if i > margin {
			return false
		}
		if margin >= len(nums) {
			return true
		}
		if next := i + nums[i]; next > margin {
			margin = next
		}
	}
	return true
}

func canJump2(nums []int) int {
	end, next := 0, 0
	var count int
	for idx, num := range nums {
		if idx > end {
			panic(fmt.Errorf("invalid input not reachable idx[%d] in val %+v", idx, nums))
		}
		if curNext := idx + num; curNext > next {
			next = curNext
		}
		if idx == len(nums)-1 {
			return count
		}
		if idx == end {
			count++
			end = next
		}
	}
	panic("invalid input")
}
