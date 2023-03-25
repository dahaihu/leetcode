package timing_wheel

import (
	"container/heap"
	"container/list"
	"sync"
)

type TimingWheel struct {
	interval int64
	tickms   int64
	current  int64
	queue    *PriorityQueue
	slots    []*slot
	m        *sync.Mutex
	next     *TimingWheel
}

type slot struct {
	added  bool
	expire int64
	jobs   *list.List
}

func (s *slot) Priority() int64 {
	return s.expire
}

func (s *slot) add(e Job) {
	s.jobs.PushBack(e)
}

func New(start, tickms int64, slotNum int, queue *PriorityQueue) *TimingWheel {
	current := start - start%tickms
	slots := make([]*slot, slotNum)
	for i := 0; i < slotNum; i++ {
		slots[i] = &slot{
			expire: current + int64(i)*(tickms),
			jobs:   list.New(),
		}
	}
	return &TimingWheel{
		tickms:   tickms,
		interval: tickms * int64(slotNum),
		slots:    slots,
		current:  current,
		m:        &sync.Mutex{},
	}
}

func (t *TimingWheel) Offer(e Job) bool {
	t.m.Lock()
	defer t.m.Unlock()

	if expire := e.Priority(); expire < t.current+t.tickms {
		return false
	} else if expire < t.current+t.interval {
		index := (expire - t.current) / t.tickms
		slot := t.slots[index]
		if !slot.added {
			heap.Push(t.queue, slot)
			slot.added = true
		}
		slot.add(e)
		return true
	} else {
		if t.next == nil {
			t.next = New(t.current, t.interval, len(t.slots), t.queue)
		}
		t.next.Offer(e)
		return true
	}
}

func (t *TimingWheel) Advance(nowms int64) {
	t.m.Lock()
	defer t.m.Unlock()

	if nowms >= t.current+t.tickms {
		t.current = nowms - nowms%t.tickms
		if t.next != nil {
			t.next.Advance(nowms)
		}
	}
}
