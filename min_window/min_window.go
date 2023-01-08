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
	for i := 0; i < len(s); i++ {
		b := s[i]
		if _, needok := need[b]; !needok {
			continue
		}
		has[b]++
		if start == -1 {
			start = i
		}
		missc, missok := miss[b]
		if !missok {
			continue
		}
		if missc > 1 {
			miss[b] = missc - 1
		} else {
			delete(miss, b)
			if len(miss) != 0 {
				continue
			}
			for j := start; j <= i; j++ {
				jb := s[j]
				if _, needok := need[jb]; !needok {
					continue
				}
				hasc, needc := has[jb], need[jb]
				if hasc > needc {
					has[jb] = hasc - 1
				} else {
					has[jb] = hasc - 1
					miss[jb] = 1
					start = j + 1
					if curLen := i - j + 1; curLen < len(result) {
						result = s[j : i+1]
					}
					break
				}
			}
		}
	}
	if len(result) == len(s)+1 {
		return ""
	}
	return result
}
