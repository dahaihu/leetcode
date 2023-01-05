package min_window

func minWindow(s string, t string) string {
	var (
		result = s + " "
		start  = -1
	)
	need, miss := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(t); i++ {
		b := t[i]
		need[b]++
		miss[b]++
	}
	has := make(map[byte]int)
	nextStart := func(start, end int) int {
		for j := start; j <= end; j++ {
			jb := s[j]
			if _, ok := need[jb]; !ok {
				continue
			}
			hc, ec := has[jb], need[jb]
			if hc > ec {
				has[jb] = hc - 1
			} else {
				miss[jb] = 1
				has[jb] = hc - 1
				if curLen := end - j + 1; curLen < len(result) {
					result = s[j : end+1]
				}
				return j + 1
			}
		}
		panic("invalid input")
	}
	for i := 0; i < len(s); i++ {
		b := s[i]
		if _, needed := need[b]; !needed {
			continue
		}
		if start == -1 {
			start = i
		}
		has[b]++
		missc, missed := miss[b]
		if !missed {
			continue
		}
		if missc > 1 {
			miss[b] = missc - 1
		} else {
			delete(miss, b)
			if len(miss) == 0 {
				start = nextStart(start, i)
			}
		}
	}

	if len(result) == len(s)+1 {
		return ""
	}
	return result
}
