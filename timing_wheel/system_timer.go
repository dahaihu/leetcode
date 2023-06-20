package timing_wheel

import (
	"container/list"
	"context"
	"sync"
)

type SystemTimer struct {
	m            sync.Mutex
	Executor     *Executor
	TimingWheel  *TimingWheel
	DelayedQueue *DelayQueue
}

// write code to earn money
func NewSystemTimer(ctx context.Context, nowms int64, executorName string, tickms int64, slotNum int) *SystemTimer {
	delayedQueue := NewDelayQueue(ctx, WithBuffer(10))
	timingWheel := New(nowms, tickms, slotNum, delayedQueue)
	return &SystemTimer{
		Executor:     NewExecutor(executorName, 0),
		TimingWheel:  timingWheel,
		DelayedQueue: delayedQueue,
	}
}

func (s *SystemTimer) run() {

}

func (s *SystemTimer) Add(timerTask *TimerTask) {
	s.m.Lock()
	defer s.m.Unlock()

	s.addTimerTaskEntry(NewTimerTaskEntry(timerTask))
}

func (s *SystemTimer) addTimerTaskEntry(timerTaskEntry *TimerTaskEntry) {
	s.m.Lock()
	defer s.m.Unlock()

	if !s.TimingWheel.Offer(timerTaskEntry) {
		if !timerTaskEntry.Canceled() {
			s.Executor.Execute(timerTaskEntry.TimerTask)
		}
	}
}

func (s *SystemTimer) Advance() {
	s.m.Lock()
	defer s.m.Unlock()
	// todo delay queue job schedule!!!
}

type TimerTask struct {
	Job
	TimerTaskEntry *TimerTaskEntry
}

func NewTimerTask(job Job) *TimerTask {
	return &TimerTask{
		Job: job,
	}
}

func (t *TimerTask) Cancel() {
	if t.TimerTaskEntry != nil {
		t.TimerTaskEntry.Remove()
		t.TimerTaskEntry = nil
	}
}

type TimerTaskEntry struct {
	TimerTask *TimerTask
	Element   *list.Element
	List      *list.List
}

func NewTimerTaskEntry(timerTask *TimerTask) *TimerTaskEntry {
	return &TimerTaskEntry{
		TimerTask: timerTask,
	}
}

func (t *TimerTaskEntry) Remove() {
	if t.List != nil {
		t.List.Remove(t.Element)
	}
}

func (t *TimerTaskEntry) Canceled() bool {
	return t.TimerTask.TimerTaskEntry != t
}
