package min_distance

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDistance(original, target string) int {
	// new and init mark
	mark := make([][]int, len(target)+1)
	for i := 0; i <= len(target); i++ {
		mark[i] = make([]int, len(original)+1)
	}
	mark[0][0] = 0
	for i := 1; i <= len(target); i++ {
		mark[i][0] = i
	}
	for i := 1; i <= len(original); i++ {
		mark[0][i] = i
	}

	for i := 1; i <= len(target); i++ {
		for j := 1; j <= len(original); j++ {
			mark[i][j] = min(
				min(mark[i][j-1], mark[i-1][j-1]),
				mark[i-1][j],
			) + 1
			if target[i-1] == original[j-1] && mark[i-1][j-1] < mark[i][j] {
				mark[i][j] = mark[i-1][j-1]
			}
		}
	}
	return mark[len(target)][len(original)]
}
