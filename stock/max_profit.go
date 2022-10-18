package stock

func maxProfit(k int, prices []int) int {
	mark := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		mark[i] = make([]int, len(prices))
	}

	for i := 1; i <= k; i++ {
		pre := -prices[0]
		for j := 1; j <= len(prices)-1; j++ {
			mark[i][j] = mark[i][j-1]
			if cur := -prices[j] + mark[i-1][j-1]; cur > pre {
				pre = cur
			}
			if profit := pre + prices[j]; profit > mark[i][j] {
				mark[i][j] = profit
			}
		}
	}

	return mark[k][len(prices)-1]
}
