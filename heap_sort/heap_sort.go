package heapsort

type queue []int

func (q *queue) down(idx int) {
	for idx < q.len() {
		left := 2*idx + 1
		if left >= q.len() {
			return
		}
		minChild := left
		if right := left + 1; right < q.len() &&
			q.indexVal(right) < q.indexVal(left) {
			minChild = right
		}
		if q.indexVal(idx) <= q.indexVal(minChild) {
			return
		}
		q.swap(idx, minChild)
		idx = minChild
	}
}

func (q *queue) len() int {
	return len(*q)
}

func (q *queue) indexVal(i int) int {
	return (*q)[i]
}

func (q *queue) indexSet(idx int, val int) {
	(*q)[idx] = val
}

func (q *queue) swap(i, j int) {
	(*q)[i], (*q)[j] = (*q)[j], (*q)[i]
}

func newQuque(nums []int) *queue {
	q := &queue{}
	for _, num := range nums {
		*q = append(*q, num)
	}
	q.init()
	return q
}

func (q *queue) init() {
	for i := (q.len() - 2) / 2; i >= 0; i-- {
		q.down(i)
	}
}

func (q *queue) push(num int) bool {
	if num <= q.indexVal(0) {
		return false
	}
	q.indexSet(0, num)
	q.down(0)
	return true
}

func down(nums []int, idx int, margin int) {
	for idx < margin {
		left := 2*idx + 1
		if left >= margin {
			return
		}
		minChild := left
		right := left + 1
		if right < margin && nums[right] > nums[minChild] {
			minChild = right
		}
		if nums[idx] >= nums[minChild] {
			return
		}
		nums[idx], nums[minChild] = nums[minChild], nums[idx]
		idx = minChild
	}
}

func heapSort(nums []int) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		down(nums, i, len(nums))
	}
	for i := len(nums) - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		down(nums, 0, i)
	}
}
