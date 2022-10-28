package next_number

import (
	"sort"
)

// low...high
func int2array(num int) []int {
	var items []int
	for num != 0 {
		items = append(items, num%10)
		num = num / 10
	}
	return items
}

func nextNumber(candidates []int, target int) int {
	sort.Ints(candidates)
	minCandidate, maxCandidate := candidates[0], candidates[len(candidates)-1]
	maxCandidateIdx := len(candidates) - 1
	targetItems := int2array(target)
	var resultIndex []int
	for i := len(targetItems) - 1; i >= 0; i-- {
		if num := targetItems[i]; num < minCandidate {
			preIdx := -1
			for j := len(resultIndex) - 1; j >= 0; j-- {
				if resultIndex[j] > 0 {
					preIdx = j
					break
				}
			}
			if preIdx == -1 {
				resultIndex = make([]int, len(targetItems)-1)
				for j := 0; j < len(targetItems)-1; j++ {
					resultIndex[j] = maxCandidateIdx
				}
			} else {
				resultIndex[preIdx] = resultIndex[preIdx] - 1
				for j := preIdx + 1; j < len(resultIndex); j++ {
					resultIndex[j] = maxCandidateIdx
				}
				for j := i; j >= 0; j-- {
					resultIndex = append(resultIndex, maxCandidateIdx)
				}
			}
			break
		} else if num > maxCandidate {
			for j := i; j >= 0; j-- {
				resultIndex = append(resultIndex, maxCandidateIdx)
			}
			break
		} else {
			idx := sort.Search(len(candidates), func(j int) bool { return num <= candidates[j] })
			if candidates[idx] == num {
				resultIndex = append(resultIndex, idx)
				continue
			}
			resultIndex = append(resultIndex, idx-1)
			for j := i - 1; j >= 0; j-- {
				resultIndex = append(resultIndex, len(candidates)-1)
			}
			break
		}
	}
	var res int
	for i := 0; i < len(resultIndex); i++ {
		res = res*10 + candidates[resultIndex[i]]
	}
	return res
}
