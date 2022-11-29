package can_partition

func canPartition(nums []int) bool {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}

	mark := make([][]bool, sum/2+1)
	for i := 0; i < sum/2+1; i++ {
		mark[i] = make([]bool, len(nums)+1)
	}
	mark[0][0] = true

	for i := 1; i < sum/2+1; i++ {
		for j := 1; j < len(nums)+1; j++ {
			if i < nums[j-1] {
				break
			}
			// candidate: nums[j-1]
			// target: i
			mark[i][j] = mark[i-nums[j-1]][j-1]
			if mark[i][j] {
				for jj := j + 1; jj < len(nums)+1; jj++ {
					mark[i][jj] = true
				}
				break
			}
		}
	}
	return mark[sum/2][len(nums)]
}
