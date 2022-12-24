package subarray_sum

func subarraySum(nums []int, target int) [][]int {
	mark := make(map[int][]int)
	mark[0] = []int{-1}
	var (
		result [][]int
		sum    int
	)
	for idx, num := range nums {
		sum += num
		if suiteds, ok := mark[sum-target]; ok {
			for _, suited := range suiteds {
				result = append(result, nums[suited+1:idx+1])
			}
		}
		mark[sum] = append(mark[sum], idx)
	}
	return result
}
