package scheduler

import (
	"container/heap"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

var token = struct{}{}

type Scheduler struct {
	priorityQueue PriorityQueue
	now           func() int64
	resume        chan struct{}
	closeChan     chan struct{}
	closedChan    chan struct{}
	concurrency   chan struct{}
	duration      time.Duration
	mutex         sync.Mutex
	waitingCount  atomic.Int64
}

func NewSchduler(concurrency int, duration time.Duration, now func() int64) *Scheduler {
	s := &Scheduler{
		concurrency: make(chan struct{}, concurrency),
		closedChan:  make(chan struct{}),
		closeChan:   make(chan struct{}),
		resume:      make(chan struct{}),
		duration:    duration,
		now:         now}
	go s.dispatch()
	return s
}

func (s *Scheduler) Push(job IJob) (jw *JobWrapper) {
	// check scheduler stopped
	select {
	case <-s.closedChan:
		panic("scheduler has been stopped")
	default:
	}

	jw = NewJobWrapper(job)

	s.mutex.Lock()
	Push(&s.priorityQueue, jw)
	s.mutex.Unlock()

	if jw.Index() == 0 {
		select {
		case s.resume <- token:
		default:
		}
	}
	return jw
}

func (s *Scheduler) Stop(f func(IJob)) {
	close(s.closeChan)
	<-s.closedChan

	if f != nil {
		for s.priorityQueue.Len() != 0 {
			jobWrapper := Pop(&s.priorityQueue).(*JobWrapper)
			f(jobWrapper.job)
		}
	}
}

func (s *Scheduler) Init(jobs ...IJob) {
	for _, job := range jobs {
		s.Push(job)
	}
}

func (s *Scheduler) wait(gap int64) (shouldContinue bool) {
	t := time.NewTimer(s.duration * time.Duration(gap))
	defer t.Stop()

	select {
	case <-t.C:
	case <-s.resume:
	case <-s.closeChan:
		return false
	}
	return true
}

func (s *Scheduler) dispatch() {
	defer func() {
		close(s.closedChan)
	}()

	for {
		jw := s.FirstJob()
		if jw == nil {
			select {
			case <-s.resume:
				continue
			case <-s.closeChan:
				return
			}
		}
		if removed := jw.Removed(); removed {
			continue
		}
		now := s.now()
		gap := jw.job.ExecuteTime() - now
		log.Printf("now %d, executeTime %d,  encounter gap %+v", now, jw.job.ExecuteTime(), gap)
		if gap <= 0 {
			if removed := s.pop(jw); !removed {
				s.run(jw.job)
			}
			continue
		}
		if shouldContinue := s.wait(gap); !shouldContinue {
			break
		}
	}
}

func (s *Scheduler) run(job IJob) {
	s.waitingCount.Add(1)
	defer s.waitingCount.Add(-1)

	s.concurrency <- token
	go func() {
		defer func() { <-s.concurrency }()
		job.Do()
	}()
}

func (s *Scheduler) pop(wrapped *JobWrapper) (deleted bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	heap.Remove(&s.priorityQueue, wrapped.Index())
	return wrapped.Removed()
}

func (s *Scheduler) FirstJob() *JobWrapper {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	job := s.priorityQueue.MaxPriority()
	if job == nil {
		return nil
	}
	return job.(*JobWrapper)
}

func (s *Scheduler) PendingJob() int64 {
	return s.waitingCount.Load()
}
