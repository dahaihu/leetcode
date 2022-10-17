package can_jump

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
