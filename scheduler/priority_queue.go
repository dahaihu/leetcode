package scheduler

import "container/heap"

type IElement interface {
	Index() int
	SetIndex(int)
	Priority() int64
}

type PriorityQueue struct {
	Jobs []IElement
}

func (p *PriorityQueue) Len() int {
	return len(p.Jobs)
}

func (p *PriorityQueue) Swap(i, j int) {
	p.Jobs[i], p.Jobs[j] = p.Jobs[j], p.Jobs[i]
	p.Jobs[i].SetIndex(i)
	p.Jobs[j].SetIndex(j)
}

func (p *PriorityQueue) Less(i, j int) bool {
	return p.Jobs[i].Priority() <= p.Jobs[j].Priority()
}

func (p *PriorityQueue) Push(ele interface{}) {
	p.Jobs = append(p.Jobs, ele.(IElement))
}

func (p *PriorityQueue) Pop() interface{} {
	last := p.Jobs[len(p.Jobs)-1]
	p.Jobs = p.Jobs[:len(p.Jobs)-1]
	return last
}

func (p *PriorityQueue) MaxPriority() IElement {
	if len(p.Jobs) == 0 {
		return nil
	}
	return p.Jobs[0]
}

func Push(p *PriorityQueue, ele IElement) {
	heap.Push(p, ele)
}

func Pop(p *PriorityQueue) IElement {
	return heap.Pop(p).(IElement)
}
