package leastinterval

type element struct {
	task byte
	cnt  int
}

func newElement(task byte, cnt int) *element {
	return &element{task: task, cnt: cnt}
}

type heap struct {
	elements []*element
}

func (h *heap) push(ele *element) {
	h.elements = append(h.elements, ele)
	h.up(len(h.elements) - 1)
}

func (h *heap) up(i int) {
	for i > 0 {
		p := (i - 1) / 2
		if h.elements[p].cnt >= h.elements[i].cnt {
			return
		}
		h.elements[p], h.elements[i] = h.elements[i], h.elements[p]
		i = p
	}
}

func (h *heap) pop() *element {
	if h.empty() {
		return nil
	}
	e := h.elements[0]
	h.elements = h.elements[1:]
	return e
}

func (h *heap) peek() *element {
	if h.empty() {
		return nil
	}
	return h.elements[0]
}

func (h *heap) len() int {
	return len(h.elements)
}

func (h *heap) empty() bool {
	return h.len() == 0
}

func (h *heap) reset() {
	h.elements = nil
}

func leastInterval(tasks []byte, n int) int {
	mark := make(map[byte]int)
	for _, task := range tasks {
		mark[task]++
	}
	h := new(heap)
	for task, cnt := range mark {
		h.push(newElement(task, cnt))
	}
	var result int
	for !h.empty() {
		e := h.peek()
		c := n + 1
		if e.cnt == 1 && h.len() <= c {
			result += h.len()
			h.reset()
			break
		}
		var t []*element
		for c > 0 {
			ele := h.pop()
			if ele == nil {
				result += c
				break
			}
			result += 1
			c -= 1
			if ele.cnt > 1 {
				ele.cnt -= 1
				t = append(t, ele)
			}
		}
		for _, ele := range t {
			h.push(ele)
		}
	}
	return result
}
