package timing_wheel

type Job interface {
	Priority() int64
}

type PriorityQueue struct {
	elements []Job
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{}
}

func (p *PriorityQueue) Less(i, j int) bool {
	return p.elements[i].Priority() <= p.elements[j].Priority()
}

func (p *PriorityQueue) Len() int {
	return len(p.elements)
}

func (p *PriorityQueue) Swap(i, j int) {
	p.elements[i], p.elements[j] = p.elements[j], p.elements[i]
}

func (p *PriorityQueue) Push(e interface{}) {
	p.elements = append(p.elements, e.(Job))
}

func (p *PriorityQueue) Pop() interface{} {
	last := p.elements[p.Len()-1]
	p.elements = p.elements[:p.Len()-1]
	return last
}
