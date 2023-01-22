package circuit_breaker

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

// station change
// closed when error increase to threshold open
// open timeout to half open
// halfopen when success increase threshold closed
var ErrorBreakerOpen = errors.New("circuit breaker is open")

type state uint32

const (
	closed state = iota
	open
	halfopen
)

type Breaker struct {
	errors           int
	errorThreshold   int
	successes        int
	successThreshold int
	timeout          time.Duration
	lock             sync.Mutex
	state            atomic.Value
	lastError        time.Time
}

func New(errorThreshold, successThreshold int, timeout time.Duration) *Breaker {
	var state atomic.Value
	state.Store(closed)
	return &Breaker{
		errorThreshold:   errorThreshold,
		successThreshold: successThreshold,
		timeout:          timeout,
		state:            state,
	}
}

func (b *Breaker) Run(work func() error) error {
	status := b.state.Load().(state)
	if status == open {
		return ErrorBreakerOpen
	}
	return b.doWork(status, work)
}

func (b *Breaker) doWork(state state, work func() error) error {
	var panicValue interface{}

	result := func() error {
		defer func() {
			panicValue = recover()
		}()
		return work()
	}()

	if result == nil && panicValue == nil && state == closed {
		return nil
	}

	b.processResult(result, panicValue)

	if panicValue != nil {
		panic(panicValue)
	}
	return result
}

func (b *Breaker) changeState(state state) {
	b.errors = 0
	b.successes = 0
	b.state.Store(state)
}

func (b *Breaker) closeBreaker() {
	b.changeState(closed)
}

func (b *Breaker) openBreaker() {
	b.changeState(open)
	go func() {
		time.Sleep(b.timeout)
		b.lock.Lock()
		defer b.lock.Unlock()
		b.changeState(halfopen)
	}()
}

func (b *Breaker) processResult(result error, panicValue interface{}) {
	b.lock.Lock()
	defer b.lock.Unlock()

	if result == nil && panicValue == nil {
		if state := b.state.Load().(state); state == halfopen {
			b.successes++
			if b.successes == b.successThreshold {
				b.closeBreaker()
			}
		}
	} else {
		switch state := b.state.Load().(state); state {
		case closed:
			if expire := b.lastError.Add(b.timeout); time.Now().After(expire) {
				b.errors = 1
				b.lastError = time.Now()
				return
			}
			b.errors++
			if b.errors == b.errorThreshold {
				b.openBreaker()
			} else {
				b.lastError = time.Now()
			}
		case halfopen:
			b.openBreaker()
		}
	}
}
