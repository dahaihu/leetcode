package subarray_sum

func subarraySum(nums []int, target int) [][]int {
	mark := make(map[int][]int)
	// cur - pre = target -> pre = cur - target
	mark[0] = []int{-1}
	var (
		sum    int
		result [][]int
	)
	for idx, num := range nums {
		sum += num
		if preEnds := mark[sum-target]; len(preEnds) > 0 {
			for _, preEnd := range preEnds {
				result = append(result, nums[preEnd+1:idx+1])
			}
		}
		mark[sum] = append(mark[sum], idx)
	}
	return result
}
