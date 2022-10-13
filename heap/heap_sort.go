package heap

type HeapInt []int

func (h *HeapInt) Less(i, j int) bool {
	return (*h)[i] <= (*h)[j]
}

func (h *HeapInt) Len() int {
	return len(*h)
}

func (h *HeapInt) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *HeapInt) Push(ele interface{}) {
	*h = append(*h, ele.(int))
}

func (h *HeapInt) Pop() interface{} {
	x := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return x
}
