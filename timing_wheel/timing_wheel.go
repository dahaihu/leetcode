package timing_wheel

import (
	"container/list"
	"sync"
	"sync/atomic"
)

type TimingWheel struct {
	tickms   int64
	interval int64
	slots    slots
	current  int64
	queue    *DelayQueue
	m        *sync.Mutex
	next     *TimingWheel
}

type Job interface {
	Priority() int64
	Do() func()
	Key() string
}

type slot struct {
	m      sync.Mutex
	expire int64
	jobs   *list.List
}

func (s *slot) ExecuteTime() int64 {
	return s.expire
}

func (s *slot) push(entry *TimerTaskEntry) {
	s.m.Lock()
	defer s.m.Unlock()

	entry.Element = s.jobs.PushBack(entry)
	entry.List = s.jobs
}

func (s *slot) setExpire(expire int64) bool {
	return atomic.CompareAndSwapInt64(&s.expire, s.expire, expire)
}

func (s *slot) resetExpire() {
	atomic.StoreInt64(&s.expire, -1)
}

func (s *slot) flush(f func(e Job)) {
	s.m.Lock()
	defer s.m.Unlock()

	for e := s.jobs.Front(); e != nil; e = s.jobs.Front() {
		timerTaskEntry := e.Value.(*TimerTaskEntry)
		f(timerTaskEntry.TimerTask.Job)
		s.jobs.Remove(e)
	}
	s.resetExpire()
}

type slots []*slot

func New(start, tickms int64, slotNum int, queue *DelayQueue) *TimingWheel {
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
		queue:    queue,
		next:     nil,
	}
}

func (t *TimingWheel) Offer(task *TimerTaskEntry) bool {
	t.m.Lock()
	defer t.m.Unlock()

	if expire := task.TimerTask.Priority(); expire < t.current+t.tickms {
		return false
	} else if expire < t.current+t.interval {
		virtualIndex := expire / t.tickms
		slot := t.slots[virtualIndex%int64(len(t.slots))]
		slot.push(task)
		if slot.setExpire(virtualIndex * t.tickms) {
			t.queue.Offer(slot)
		}
		return true
	} else {
		if t.next == nil {
			t.next = New(t.current, t.interval, len(t.slots), t.queue)
		}
		t.next.Offer(task)
		return true
	}
}

func (t *TimingWheel) Poll() {
	t.m.Lock()
	defer t.m.Unlock()

}

func (t *TimingWheel) Advance(nowms int64) {
	if nowms >= t.current+t.tickms {
		t.current = nowms - nowms%t.tickms
		if t.next != nil {
			t.next.Advance(nowms)
		}
	}
}
