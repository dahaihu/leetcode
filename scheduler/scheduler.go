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
	waitingCount  int64
}

// NewSchuduler create a scheduler to schedule job
func NewSchduler(concurrency int, duration time.Duration, now func() int64) *Scheduler {
	var concurrencyCh chan struct{}
	if concurrency > 0 {
		concurrencyCh = make(chan struct{}, concurrency)
	}
	s := &Scheduler{
		concurrency: concurrencyCh,
		closedChan:  make(chan struct{}),
		closeChan:   make(chan struct{}),
		resume:      make(chan struct{}),
		duration:    duration,
		now:         now,
	}
	go s.dispatch()
	return s
}

// Push pushes the job to the scheduler. when sheduler stopped, it panics
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

	// signal to dispatch, let dispatch know that new max priority job arrives
	if jw.Index() == 0 {
		select {
		case s.resume <- token:
		default:
		}
	}

	return jw
}

// Stop stops the scheduler, use f to process the unprocessed job
func (s *Scheduler) Stop(f func(IJob)) {
	// signal to dispatch, let dispatch know that scheduler stopped
	close(s.closeChan)
	// wait for dispatch goroutine terminated
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

// PendingJobNum returns the job number which should execute now but block for concurrency limit
func (s *Scheduler) PendingJobNum() int64 {
	return atomic.LoadInt64(&s.waitingCount)
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
		jw := s.nextJob()
		if jw == nil {
			select {
			case <-s.resume:
				continue
			case <-s.closeChan:
				return
			}
		}
		// job was removed, just pop job from heap is ok
		if removed := jw.Removed(); removed {
			s.pop(jw)
			continue
		}
		now := s.now()
		gap := jw.job.ExecuteTime() - now
		log.Printf("now %d, executeTime %d,  encounter gap %+v", now, jw.job.ExecuteTime(), gap)
		if gap <= 0 {
			// job should do now
			if removed := s.pop(jw); !removed {
				s.run(jw.job)
			}
			continue
		} else {
			// gap > 0, wait for gap elapse or new max priority job arrives or scheduler stopped
			if shouldContinue := s.wait(gap); !shouldContinue {
				break
			}
		}
	}
}

func (s *Scheduler) run(job IJob) {
	atomic.AddInt64(&s.waitingCount, 1)
	defer atomic.AddInt64(&s.waitingCount, -1)

	if s.concurrency != nil {
		s.concurrency <- token
	}

	go func() {
		defer func() {
			if s.concurrency != nil {
				<-s.concurrency
			}
		}()

		job.Do()
	}()
}

// pop remove wrapped job from priority queue
func (s *Scheduler) pop(wrapped *JobWrapper) (deleted bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	heap.Remove(&s.priorityQueue, wrapped.Index())
	return wrapped.Removed()
}

// nextJob returns the max priority job, if empty returns nil
func (s *Scheduler) nextJob() *JobWrapper {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	job := s.priorityQueue.MaxPriority()
	if job == nil {
		return nil
	}
	return job.(*JobWrapper)
}
