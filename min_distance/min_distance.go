package min_distance

func minDistance(original, target string) int {
	mark := make([][]int, len(target)+1)
	for i := 0; i <= len(target); i++ {
		mark[i] = make([]int, len(original)+1)
	}
	for i := 0; i <= len(target); i++ {
		mark[i][0] = i
	}
	for i := 1; i <= len(original); i++ {
		mark[0][i] = i
	}
	for i := 1; i <= len(target); i++ {
		for j := 1; j <= len(original); j++ {
			// 增加
			cur := mark[i-1][j] + 1
			// 删除
			if deleted := mark[i][j-1] + 1; deleted < cur {
				cur = deleted
			}
			// 修改
			if target[i-1] == original[j-1] {
				if mark[i-1][j-1] < cur {
					cur = mark[i-1][j-1]
				}
			} else {
				if updated := mark[i-1][j-1] + 1; updated < cur {
					cur = updated
				}
			}

			mark[i][j] = cur
		}
	}
	return mark[len(target)][len(original)]
}
