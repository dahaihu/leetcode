package next_number

import (
	"sort"
)

// []int{low...high}
func int2array(num int) []int {
	var items []int
	for num != 0 {
		items = append(items, num%10)
		num = num / 10
	}
	return items
}

func nextNumber(candidates []int, target int) int {
	sort.Slice(candidates, func(i, j int) bool { return candidates[i] <= candidates[j] })
	maxCandidateIndex := len(candidates) - 1
	targetItems := int2array(target)
	var resultIndexs []int
	for i := len(targetItems) - 1; i >= 0; i-- {
		item := targetItems[i]
		itemIdx := sort.Search(len(candidates), func(j int) bool { return item <= candidates[j] })
		if itemIdx < len(candidates) && candidates[itemIdx] == item {
			resultIndexs = append(resultIndexs, itemIdx)
			continue
		}
		if itemIdx == 0 {
			preSuitedIdx := -1
			for j := len(resultIndexs) - 1; j >= 0; j-- {
				if resultIndexs[j] > 0 {
					preSuitedIdx = j
					break
				}
			}
			if preSuitedIdx != -1 {
				resultIndexs[preSuitedIdx] = resultIndexs[preSuitedIdx] - 1
				for j := preSuitedIdx; j < len(resultIndexs); j++ {
					resultIndexs[j] = maxCandidateIndex
				}
				for j := i; j >= 0; j-- {
					resultIndexs = append(resultIndexs, maxCandidateIndex)
				}
			} else {
				resultIndexs = make([]int, len(targetItems)-1)
				for j := 0; j < len(resultIndexs); j++ {
					resultIndexs[j] = maxCandidateIndex
				}
			}
		} else {
			resultIndexs = append(resultIndexs, itemIdx-1)
			for j := i - 1; j >= 0; j-- {
				resultIndexs = append(resultIndexs, maxCandidateIndex)
			}
		}
		break
	}
	var result int
	for _, index := range resultIndexs {
		result = 10*result + candidates[index]
	}
	return result
}
