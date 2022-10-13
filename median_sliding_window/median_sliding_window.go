package median_sliding_window

import "sort"

type window []int

func (w window) remove(ele int) {
	elePos := sort.Search(len(w), func(i int) bool { return ele <= w[i] })
	copy(w[elePos:], w[elePos+1:])
}

func (w window) add(ele int) {
	elePos := sort.Search(len(w)-1, func(i int) bool { return ele <= w[i] })
	copy(w[elePos+1:], w[elePos:])
	w[elePos] = ele
}

func (w window) median() float64 {
	if len(w)%2 == 1 {
		return float64(w[len(w)/2])
	}
	return float64(w[len(w)/2-1]+w[len(w)/2]) / 2
}

func medianSlidingWindow(nums []int, k int) []float64 {
	queue := window(append([]int{}, nums[:k]...))
	sort.Slice(queue, func(i, j int) bool { return queue[i] <= queue[j] })
	var result []float64
	result = append(result, queue.median())
	for i := k; i < len(nums); i++ {
		queue.remove(nums[i-k])
		queue.add(nums[i])
		result = append(result, queue.median())
	}
	return result
}
