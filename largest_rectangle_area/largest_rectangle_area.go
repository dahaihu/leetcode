package largest_rectangle_area

type element struct {
	idx    int
	height int
}

func (e *element) area(idx int) int {
	return (idx - e.idx) * e.height
}

type queue []*element

func (q *queue) tailHeight() int {
	return (*q)[q.len()-1].height
}

func (q *queue) pop() *element {
	last := (*q)[q.len()-1]
	*q = (*q)[:q.len()-1]
	return last
}

func (q *queue) empty() bool {
	if q == nil {
		return true
	}
	return len(*q) == 0
}

func (q *queue) len() int {
	if q == nil {
		return 0
	}
	return len(*q)
}

func (q *queue) append(idx, height int) {
	*q = append(*q, &element{idx: idx, height: height})
}

func largestRectangleArea(heights []int) int {
	var largestArea int
	q := new(queue)
	for idx, height := range heights {
		if q.empty() || q.tailHeight() < height {
			q.append(idx, height)
			continue
		}
		var last *element
		for !q.empty() && q.tailHeight() >= height {
			last = q.pop()
			if area := last.area(idx); area > largestArea {
				largestArea = area
			}
		}
		q.append(last.idx, height)
	}
	margin := len(heights)
	for q.len() != 0 {
		last := q.pop()
		if area := last.area(margin); area > largestArea {
			largestArea = area
		}
	}
	return largestArea
}
