package max_sliding_window

type element struct {
	val int
	idx int
}

func newElement(array []int, idx int) *element {
	return &element{val: array[idx], idx: idx}
}

func (e *element) Val() int {
	return e.val
}

func (e *element) Index() int {
	return e.idx
}

func maxSlidingWindow(nums []int, k int) []int {
	queue := []*element{newElement(nums, 0)}
	for i := 1; i < k; i++ {
		for len(queue) != 0 && queue[len(queue)-1].Val() <= nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, newElement(nums, i))
	}
	result := []int{queue[0].Val()}
	for i := k; i < len(nums); i++ {
		for len(queue) != 0 && queue[0].Index() <= i-k {
			queue = queue[1:]
		}
		for len(queue) != 0 && queue[len(queue)-1].Val() <= nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, newElement(nums, i))
		result = append(result, queue[0].Val())
	}
	return result
}
