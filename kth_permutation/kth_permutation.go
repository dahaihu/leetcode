package kthpermulation

import "strconv"

func getPermutation(n int, k int) string {
	mark := make([]int, n)
	mark[0] = 1
	for i := 1; i < n; i++ {
		mark[i] = mark[i-1] * i
	}
	if k > n*mark[n-1] {
		panic("k overflow n!")
	}
	items := make([]int, n)
	for i := 1; i <= n; i++ {
		items[i-1] = i
	}
	var (
		result, idx int
		remainder   int = k
	)
	for {
		idx, remainder = k/mark[len(items)-1], k%mark[len(items)-1]
		if remainder == 0 {
			result = 10*result + items[idx-1]
			items = pop(items, idx-1)
			for j := len(items) - 1; j >= 0; j-- {
				result = 10*result + items[j]
			}
			break
		}
		result = 10*result + items[idx]
		items = pop(items, idx)
	}
	return strconv.Itoa(result)
}

func pop(items []int, idx int) []int {
	copy(items[idx:], items[idx+1:])
	return items[:len(items)-1]
}
