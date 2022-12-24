package sum_combination

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var dfs func(int, int, []int)
	dfs = func(idx int, curTarget int, temp []int) {
		if curTarget == 0 {
			result = append(result, temp)
			return
		}
		if idx == len(candidates) || curTarget < 0 {
			return
		}
		dfs(idx+1, curTarget, temp)

		nextTemp := make([]int, len(temp)+1)
		copy(nextTemp[:len(temp)], temp)
		nextTemp[len(temp)] = candidates[idx]
		dfs(idx, curTarget-candidates[idx], nextTemp)
	}
	dfs(0, target, []int{})
	return result
}

func combinationSum2(candidates []int, target int) [][]int {
	var result [][]int
	var dfs func(int, int, []int)
	dfs = func(idx int, curTarget int, inter []int) {
		if curTarget == 0 {
			result = append(result, inter)
			return
		}
		for curIdx := idx; curIdx < len(candidates); curIdx++ {
			if curTarget < candidates[curIdx] {
				break
			}
			if curIdx > idx && candidates[curIdx] == candidates[curIdx-1] {
				continue
			}
			nextInter := make([]int, len(inter)+1)
			copy(nextInter[:len(inter)], inter)
			nextInter[len(inter)] = candidates[curIdx]
			dfs(curIdx+1, curTarget-candidates[curIdx], nextInter)
		}
	}
	sort.Ints(candidates)
	dfs(0, target, []int{})
	return result
}
