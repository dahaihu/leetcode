package min_window

type mark struct {
	has    map[byte]int // has已有的数量
	need   map[byte]int // 需要的元素数量
	target map[byte]int // 目标的的元素数量
}

func newMark(s string) *mark {
	m := new(mark)
	m.need = make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m.need[s[i]]++
	}
	m.target = make(map[byte]int)
	for k, c := range m.need {
		m.target[k] = c
	}
	m.has = make(map[byte]int)
	return m
}

func (m *mark) add(b byte) (added bool, matched bool) {
	if hasc, has := m.has[b]; has {
		m.has[b] = hasc + 1
	}
	needc, ok := m.need[b]
	if !ok {
		return false, false
	}
	if _, has := m.has[b]; !has {
		m.has[b] = 1
	}
	if needc > 1 {
		m.need[b] = needc - 1
		return true, false
	}
	delete(m.need, b)
	return true, len(m.need) == 0
}

func (m *mark) remove(b byte) (ok bool, matched bool) {
	hasc, has := m.has[b]
	if !has {
		return false, true
	}
	m.has[b] = hasc - 1
	if hasc-1 < m.target[b] {
		m.need[b] = 1
		return true, false
	}
	return false, true
}

func (m *mark) unnecessary(b byte) bool {
	if _, ok := m.target[b]; !ok {
		return true
	}
	return false
}

func minWindow(s string, target string) string {
	m := newMark(target)
	var (
		pre    int    = -1
		result string = s + " "
	)
	for idx := 0; idx < len(s); idx++ {
		added, matched := m.add(s[idx])
		if !added {
			continue
		}
		if matched {
			if pre == -1 {
				return s[idx : idx+1]
			}
			if length := idx - pre + 1; length < len(result) {
				result = s[pre : idx+1]
			}
			for pre < idx {
				ok, matched := m.remove(s[pre])
				if matched {
					if length := (idx - (pre + 1) + 1); length < len(result) {
						result = s[pre+1 : idx+1]
					}
				}
				pre++
				if ok {
					break
				}
			}
			for pre < idx && m.unnecessary(s[pre]) {
				pre++
			}
			if idx == len(s) {
				break
			}
		}
		if pre == -1 {
			pre = idx
		}
	}
	if len(result) == len(s)+1 {
		return ""
	}
	return result
}
