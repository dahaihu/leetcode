package tunny

import (
	"sync"
	"sync/atomic"
)

type Pool struct {
	requestChan  chan *workRequest
	workers      []*workerWrapper
	workerCreate func() Worker
	queuedJobs   int64
	mutex        sync.Mutex
}

func New(numWorkers int, workerCreate func() Worker) *Pool {
	p := &Pool{
		requestChan:  make(chan *workRequest),
		workerCreate: workerCreate,
	}
	p.SetSize(numWorkers)
	return p
}

func NewFunc(numWorkers int, f func(interface{}) interface{}) *Pool {
	return New(numWorkers, func() Worker { return &closureWorker{processor: f} })
}

func (p *Pool) Process(payload interface{}) interface{} {
	atomic.AddInt64(&p.queuedJobs, 1)
	defer atomic.AddInt64(&p.queuedJobs, -1)

	request, ok := <-p.requestChan
	if !ok {
		panic("pool closed")
	}
	request.payloadChan <- payload

	ret, ok := <-request.retChan
	if !ok {
		panic("worker closed")
	}
	return ret
}

func (p *Pool) SetSize(size int) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	length := len(p.workers)
	switch {
	case length == size:
		return
	case size > length:
		for i := length; i < size; i++ {
			p.workers = append(p.workers, newWorkerWrapper(p.requestChan, p.workerCreate()))
		}
	case size < length:
		for i := size; i < length; i++ {
			p.workers[i].stop()
		}
		for i := size; i < length; i++ {
			p.workers[i].join()
			p.workers[i] = nil
		}
		p.workers = p.workers[:size]
	}
}

func (p *Pool) QueuedLength() int64 {
	return atomic.LoadInt64(&p.queuedJobs)
}

func (p *Pool) Close() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.SetSize(0)
	close(p.requestChan)
}

type closureWorker struct {
	processor func(interface{}) interface{}
}

func (c *closureWorker) Process(payload interface{}) interface{} {
	return c.processor(payload)
}

func (c *closureWorker) Interrupt()       {}
func (c *closureWorker) BlockUntilReady() {}
func (c *closureWorker) Terminate()       {}
