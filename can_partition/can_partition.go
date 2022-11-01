package can_partition

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	mark := make([][]bool, target)
	for i := 0; i < target; i++ {
		mark[i] = make([]bool, len(nums))
	}
	for i := 0; i < target; i++ {
		for j := 0; j < len(nums); j++ {
			cur := i + 1
			if cur == nums[j] {
				mark[i][j] = true
				continue
			}
			if j >= 1 {
				mark[i][j] = mark[i][j-1]
				if i >= nums[j] {
					mark[i][j] = mark[i][j] || mark[cur-nums[j]][j-1]
				}
			}
		}
	}
	return mark[target-1][len(nums)-1]
}
