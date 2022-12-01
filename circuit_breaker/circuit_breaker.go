package circuit_breaker

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

var ErrorBreakerOpen = errors.New("circuit breaker is open")

const (
	closed uint32 = iota
	open
	halfOpen
)

type Breaker struct {
	errors           int
	errorThreshold   int
	successes        int
	successThreshold int
	timeout          time.Duration
	lock             sync.Mutex
	state            uint32
	lastError        time.Time
}

func New(errorThreshold, successThreshold int, timeout time.Duration) *Breaker {
	return &Breaker{
		errorThreshold:   errorThreshold,
		successThreshold: successThreshold,
		timeout:          timeout,
		state:            closed,
	}
}

func (b *Breaker) Run(work func() error) error {
	status := atomic.LoadUint32(&b.state)
	if status == open {
		return ErrorBreakerOpen
	}
	return b.doWork(status, work)
}

func (b *Breaker) doWork(state uint32, work func() error) error {
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

func (b *Breaker) changeState(state uint32) {
	b.errors = 0
	b.successes = 0
	atomic.StoreUint32(&b.state, state)
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
		b.changeState(halfOpen)
	}()
}

func (b *Breaker) processResult(result error, panicValue interface{}) {
	b.lock.Lock()
	defer b.lock.Unlock()

	if result == nil && panicValue == nil {
		if b.state == halfOpen {
			b.successes++
			if b.successes == b.successThreshold {
				b.closeBreaker()
			}
		}
	} else {
		switch b.state {
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
		case halfOpen:
			b.openBreaker()
		}
	}
}
