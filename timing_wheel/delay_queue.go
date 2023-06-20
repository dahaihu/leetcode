package timing_wheel

import (
	"container/heap"
	"context"
	"errors"
	"sync"
	"time"
)

type Element interface {
	ExecuteTime() int64
}

type element struct {
	Element

	index int
}

type DelayQueue struct {
	ctx      context.Context
	cancel   context.CancelFunc
	elements []*element
	m        sync.Mutex
	C        chan Element
	resume   chan struct{}
	closed   chan struct{}
	now      func() int64
	flush    func(Element)
}

func now() int64 {
	return time.Now().UnixMilli()
}

type Option func(*DelayQueue)

func WithFlush(f func(Element)) Option {
	return func(queue *DelayQueue) {
		queue.flush = f
	}
}

func WithNow(now func() int64) Option {
	return func(queue *DelayQueue) {
		queue.now = now
	}
}

func WithBuffer(bufferSize int) Option {
	return func(queue *DelayQueue) {
		queue.C = make(chan Element, bufferSize)
	}
}

func NewDelayQueue(ctx context.Context, options ...Option) *DelayQueue {
	ctx, cancel := context.WithCancel(ctx)
	queue := &DelayQueue{
		ctx:    ctx,
		cancel: cancel,
		C:      make(chan Element),
		now:    now,
		resume: make(chan struct{}),
		closed: make(chan struct{}),
	}

	for _, op := range options {
		op(queue)
	}

	go queue.start()

	return queue
}

func (p *DelayQueue) Less(i, j int) bool {
	return p.elements[i].ExecuteTime() <= p.elements[j].ExecuteTime()
}

func (p *DelayQueue) Len() int {
	return len(p.elements)
}

func (p *DelayQueue) Swap(i, j int) {
	p.elements[i], p.elements[j] = p.elements[j], p.elements[i]
	p.elements[i].index = i
	p.elements[j].index = j
}

func (p *DelayQueue) Push(e interface{}) {
	p.elements = append(p.elements, e.(*element))
}

func (p *DelayQueue) Pop() interface{} {
	last := p.elements[p.Len()-1]
	p.elements = p.elements[:p.Len()-1]
	return last
}

func (p *DelayQueue) peek() *element {
	if p.Len() == 0 {
		return nil
	}
	return p.elements[0]
}

var ErrClosed = errors.New("delay queue closed")

func (p *DelayQueue) Offer(j Element) error {
	select {
	case <-p.ctx.Done():
		return ErrClosed
	default:
	}

	element := &element{Element: j}

	p.m.Lock()
	heap.Push(p, element)
	p.m.Unlock()

	if element.index == 0 {
		select {
		case p.resume <- struct{}{}:
		default:
		}
	}
	return nil
}

func (p *DelayQueue) Stop() {
	p.cancel()
}

func (p *DelayQueue) Wait() {
	<-p.closed
}

func (p *DelayQueue) start() {
	for {
		p.m.Lock()
		first := p.peek()
		p.m.Unlock()

		if first == nil {
			select {
			case <-p.ctx.Done():
				goto exit
			case <-p.resume:
				continue
			}
		}

		tickTime := p.now()
		gap := first.ExecuteTime() - tickTime
		if gap <= 0 {
			p.m.Lock()
			first := heap.Pop(p).(*element)
			p.m.Unlock()

			p.C <- first.Element
			continue
		}
		tick := time.NewTicker(time.Millisecond * time.Duration(gap))
		select {
		case <-p.ctx.Done():
			goto exit
		case <-tick.C:
			continue
		case <-p.resume:
			continue
		}
	}
exit:
	if p.flush != nil {
		for p.Len() != 0 {
			element := heap.Pop(p).(*element)
			p.flush(element.Element)
		}
	}
	close(p.closed)
}
