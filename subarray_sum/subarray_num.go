package subarray_sum

func subarraySum(nums []int, target int) [][]int {
	mark := make(map[int][]int)
	mark[0] = []int{-1}
	var (
		sum int
		out [][]int
	)
	for idx, num := range nums {
		sum += num
		if preEnds := mark[sum-target]; len(preEnds) > 0 {
			for _, preEnd := range preEnds {
				out = append(out, nums[preEnd+1:idx+1])
			}
		}
		mark[sum] = append(mark[sum], idx)
	}
	return out
}
