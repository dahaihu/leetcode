package subarray_sum

func subarraySum(nums []int, target int) [][]int {
	mark := make(map[int][]int)
	mark[0] = []int{-1}
	var (
		sum    int
		result [][]int
	)
	for idx, num := range nums {
		sum += num
		if ends := mark[sum-target]; len(ends) > 0 {
			for _, end := range ends {
				result = append(result, nums[end+1:idx+1])
			}
		}
		mark[sum] = append(mark[sum], idx)
	}
	return result
}
