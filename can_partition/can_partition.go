package can_partition

import "sort"

func canPartition(nums []int) bool {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}

	// mark := make([][]bool, sum/2+1)
	// for i := 0; i < sum/2+1; i++ {
	// 	mark[i] = make([]bool, len(nums)+1)
	// }
	// mark[0][0] = true

	// for t := 1; t <= sum/2; t++ {
	// 	for i := 1; i <= len(nums); i++ {
	// 		mark[t][i] = mark[t][i-1]
	// 		if pre := t - nums[i-1]; pre >= 0 {
	// 			mark[t][i] = mark[t][i] || mark[pre][i-1]
	// 		}
	// 	}
	// }

	// return mark[sum/2][len(nums)]
	mark := make([]bool, sum/2+1)
	mark[0] = true
	for i := 1; i <= len(nums); i++ {
		for t := sum / 2; t >= 1; t-- {
			mark[t] = mark[t] || mark[t-nums[i-1]]
		}
	}
	return mark[sum/2]
}

func arrayAppend(nums [][]int, ele int) [][]int {
	result := make([][]int, 0, len(nums))
	for _, num := range nums {
		cur := append([]int{}, num...)
		cur = append(cur, ele)
		result = append(result, cur)
	}
	return result
}

func canPartitionOne(nums []int) [][]int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return nil
	}
	mark := make([][][]int, sum/2+1)
	mark[0] = [][]int{{}}
	for i := 1; i <= len(nums); i++ {
		num := nums[i-1]
		for t := sum / 2; t >= 1; t-- {
			if t < num {
				break
			}
			if pre := mark[t-num]; pre != nil {
				mark[t] = append(mark[t], arrayAppend(pre, num)...)
			}
		}
	}
	return mark[sum/2]
}

func combinations(nums []int, target int) [][]int {
	sort.Ints(nums)
	mark := make([][][]int, target+1)
	mark[0] = [][]int{{}}
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		for t := target; t >= 1; t-- {
			if t < num {
				break
			}
			if pre := mark[t-num]; pre != nil {
				mark[t] = append(mark[t], arrayAppend(pre, num)...)
			}
		}
	}
	return mark[target]
}

func combinations2(nums []int, target int) [][]int {
	sort.Ints(nums)
	mark := make([][][][]int, len(nums)+1)
	for i := 0; i <= len(nums); i++ {
		mark[i] = make([][][]int, target+1)
		mark[i][0] = [][]int{{}}
	}
	pre := 0
	for i := 1; i <= len(nums); i++ {
		num := nums[i-1]
		var start int
		if i > 1 && num == nums[i-2] {
			start = pre + 1
		} else {
			start = 1
		}
		for t := 1; t <= target; t++ {
			if t < start || t < num {
				mark[i][t] = mark[i-1][t]
				continue
			}
			mark[i][t] = append(mark[i-1][t], arrayAppend(mark[i-1][t-num], num)...)
		}
		pre = pre + num
	}
	return mark[len(nums)][target]
}
