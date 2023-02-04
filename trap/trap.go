package trap

type queue struct {
	height     []int
	indexQueue []int
}

func newQueue(height []int) *queue {
	return &queue{height: height}
}

func (q *queue) len() int {
	return len(q.indexQueue)
}

func (q *queue) empty() bool {
	return q.len() == 0
}

func (q *queue) last() (idx int, height int) {
	idx = q.indexQueue[q.len()-1]
	height = q.height[idx]
	return
}

func (q *queue) removeLast() {
	q.indexQueue = q.indexQueue[:q.len()-1]
}

func (q *queue) append(idx int) {
	q.indexQueue = append(q.indexQueue, idx)
}

func trap(height []int) int {
	var waters int
	q := newQueue(height)
	for idx, h := range height {
		if q.empty() {
			q.append(idx)
			continue
		}
		lastIdx, lastHeight := q.last()
		if lastHeight > h {
			q.append(idx)
			continue
		}
		var preHeight int
		for {
			waters += (lastHeight - preHeight) * (idx - lastIdx - 1)
			q.removeLast()
			if q.empty() {
				break
			}
			preHeight = lastHeight
			lastIdx, lastHeight = q.last()
			if lastHeight > h {
				waters += (h - preHeight) * (idx - lastIdx - 1)
				break
			}
		}
		q.append(idx)
	}
	return waters
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func trapMethod2(height []int) int {
	// 接雨水
	//
	left := make([]int, len(height))
	left[0] = height[0]
	for i := 1; i < len(height); i++ {
		left[i] = max(left[i-1], height[i])
	}

	right := make([]int, len(height))
	right[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}
	var waters int
	for i := 1; i < len(height)-1; i++ {
		cur := height[i]
		if margin := min(left[i-1], right[i+1]); margin > cur {
			waters += margin - cur
		}
	}
	return waters
}
