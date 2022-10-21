package tunny

import (
	"sync"
	"sync/atomic"
)

type Pool struct {
	queuedJobs   int64
	workers      []*workerWrapper
	reqChan      chan workRequest
	createWorker func() Worker
	mutex        sync.Mutex
}

func New(size int, factory func() Worker) *Pool {
	p := &Pool{
		queuedJobs:   0,
		workers:      nil,
		reqChan:      make(chan workRequest),
		createWorker: factory,
		mutex:        sync.Mutex{},
	}
	p.SetSize(size)
	return p
}

func NewFunc(size int, f func(interface{}) interface{}) *Pool {
	return New(size, func() Worker { return &closureWorker{processor: f} })
}

func (p *Pool) SetSize(size int) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	workerLen := len(p.workers)
	switch {
	case size == workerLen:
	case size > workerLen:
		for i := workerLen; i < size; i++ {
			p.workers = append(p.workers, newWorkerWrapper(p.reqChan, p.createWorker()))
		}
	case size < workerLen:
		for i := size; i < workerLen; i++ {
			p.workers[i].stop()
		}
		for i := size; i < workerLen; i++ {
			p.workers[i].join()
			p.workers[i] = nil
		}
		p.workers = p.workers[:size]
	}
}

func (p *Pool) Process(payload interface{}) interface{} {
	atomic.AddInt64(&p.queuedJobs, 1)
	defer atomic.AddInt64(&p.queuedJobs, -1)

	request, ok := <-p.reqChan
	if !ok {
		panic("pool closed")
	}

	request.PayloadChan <- payload

	response, ok := <-request.RetChan
	if !ok {
		panic("worker terminated")
	}
	return response
}

type closureWorker struct {
	processor func(interface{}) interface{}
}

func (c *closureWorker) Interrupt() {}

func (c *closureWorker) BlockUntilReady() {}

func (c *closureWorker) Terminate() {}

func (c *closureWorker) Process(in interface{}) interface{} {
	return c.processor(in)
}
