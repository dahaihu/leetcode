package batch

import (
	"container/list"
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

var ErrorBatchClosed = errors.New("batch closed")

type Job[K comparable, V any] interface {
	Do() error
	Key() K
	Value() V
	UpdateValue(val V) (ready bool)
}

type job[K comparable, V any] struct {
	Job[K, V]

	expire  int64
	promise *promise
}

type Promise interface {
	Error() error
}

type promise struct {
	err error
	wg  sync.WaitGroup
}

func (p *promise) Error() error {
	p.wg.Wait()
	return p.err
}

type Statistic interface {
	ScheduledNum() int64
	CompactedNum() int64
	WaitingNum() int64
}

type Batch[K comparable, V any] interface {
	Statistic

	Add(Job[K, V]) error
	AddAndWait(Job[K, V]) (Promise, error)
	StopAndWait()
}

type batch[K comparable, V any] struct {
	l          sync.Mutex
	queue      *list.List
	jobs       map[K]*list.Element
	ctx        context.Context
	cancel     context.CancelFunc
	timeWindow int64
	resume     chan struct{}
	done       chan struct{}
	sema       sema

	compacted int64
	waiting   int64
	scheduled int64
}

// New create a batch
func New[K comparable, V any](ctx context.Context, timeWindow int64, concurrency int) Batch[K, V] {
	b := new(batch[K, V])
	b.ctx, b.cancel = context.WithCancel(ctx)
	b.timeWindow = timeWindow
	b.jobs = make(map[K]*list.Element)
	b.queue = list.New()
	b.resume = make(chan struct{})
	b.done = make(chan struct{})
	b.sema = newSema(concurrency)
	go b.start()
	return b
}

func (b *batch[K, V]) stop() {
	for b.queue.Len() != 0 {
		head := b.queue.Front()
		job := head.Value.(*job[K, V])
		b.do(job)
		b.queue.Remove(head)
	}
	close(b.done)
}

func (b *batch[K, V]) start() {
	var t *time.Timer
	for {
		var wrappedJob *list.Element
		b.l.Lock()
		wrappedJob = b.queue.Front()
		if wrappedJob == nil {
			b.l.Unlock()
			select {
			case <-b.ctx.Done():
				b.stop()
				return
			case <-b.resume:
				continue
			}
		}
		now := time.Now().UnixMilli()
		job := wrappedJob.Value.(*job[K, V])
		if gap := job.expire - now; gap <= 0 {
			b.queue.Remove(wrappedJob)
			delete(b.jobs, job.Key())
			t = nil
		} else {
			t = time.NewTimer(time.Millisecond * time.Duration(gap))
		}
		b.l.Unlock()

		// job expired
		if t == nil {
			b.do(job)
			continue
		}
		select {
		case <-t.C:
			continue
		case <-b.ctx.Done():
			b.stop()
			return
		}
	}
}

// AddAndWait add a job, and return a promise for waiting the job result
func (b *batch[K, V]) AddAndWait(job Job[K, V]) (promise Promise, err error) {
	return b.add(job, true)
}

// Add add a job, ignore the job result
func (b *batch[K, V]) Add(job Job[K, V]) error {
	_, err := b.add(job, false)
	return err
}

func (b *batch[K, V]) add(j Job[K, V], wait bool) (Promise, error) {
	select {
	case <-b.done:
		return nil, ErrorBatchClosed
	default:
	}
	atomic.AddInt64(&b.scheduled, 1)

	b.l.Lock()
	defer b.l.Unlock()

	wrappedJob, ok := b.jobs[j.Key()]
	if !ok {
		var p *promise
		if wait {
			p = new(promise)
			p.wg.Add(1)
		}
		job := &job[K, V]{
			Job:     j,
			expire:  b.timeWindow + time.Now().UnixMilli(),
			promise: p,
		}
		wrappedJob = b.queue.PushBack(job)
		b.jobs[j.Key()] = wrappedJob

		select {
		case b.resume <- struct{}{}:
		default:
		}

		return p, nil
	} else {
		atomic.AddInt64(&b.compacted, 1)
		oldJob := wrappedJob.Value.(*job[K, V])
		if ready := oldJob.UpdateValue(j.Value()); ready {
			b.queue.Remove(wrappedJob)
			delete(b.jobs, oldJob.Key())
			b.do(oldJob)
		}
		return oldJob.promise, nil
	}
}

func (b *batch[K, V]) do(j *job[K, V]) {
	b.sema.acquire()
	atomic.AddInt64(&b.waiting, 1)

	go func() {
		defer func() {
			b.sema.release()
			atomic.AddInt64(&b.waiting, -1)
		}()

		err := j.Job.Do()
		if j.promise != nil {
			j.promise.err = err
			j.promise.wg.Done()
		}
	}()
}

// StopAndWait stop the batch, can not add any job any more. and wait for all
// the added job to be done
func (b *batch[K, V]) StopAndWait() {
	b.cancel()

	<-b.done
}

// ScheduledNum return the number of jobs that have benn added to batch
func (b *batch[K, V]) ScheduledNum() int64 {
	return atomic.LoadInt64(&b.scheduled)
}

// CompactedNum return the number of jobs that have been compacted
func (b *batch[K, V]) CompactedNum() int64 {
	return atomic.LoadInt64(&b.compacted)
}

// WaitingNum return the number of jobs that are waiting for sema
func (b *batch[K, V]) WaitingNum() int64 {
	return atomic.LoadInt64(&b.waiting)
}
