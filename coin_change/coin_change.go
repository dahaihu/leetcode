package coin_change

import (
	"math"
	"sort"
)

func coinChange(coins []int, amount int) int {
	sort.Ints(coins)
	mark := make([]int, amount+1)
	mark[0] = 0
	for i := 1; i <= amount; i++ {
		mark[i] = -1
	}
	for i := coins[0]; i <= amount; i++ {
		cost := math.MaxInt64
		for _, coin := range coins {
			if i < coin {
				break
			}
			if curCost := mark[i-coin]; curCost != -1 && curCost < cost {
				cost = curCost
			}
		}
		if cost == math.MaxInt64 {
			continue
		}
		mark[i] = cost + 1
	}
	return mark[amount]
}
