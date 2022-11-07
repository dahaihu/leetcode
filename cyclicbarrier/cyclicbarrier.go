package cyclicbarrier

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

type round struct {
	count    int
	waitch   chan struct{}
	brokech  chan struct{}
	isBroken bool
}

type CyclicBarrier struct {
	parties int
	lock    sync.Mutex
	round   *round
}

func New(parties int) *CyclicBarrier {
	return &CyclicBarrier{
		parties: parties,
		lock:    sync.Mutex{},
		round: &round{
			waitch:  make(chan struct{}),
			brokech: make(chan struct{}),
		},
	}
}

var ErrBrokenBarrier = errors.New("cyclic barrier broken")

func (b *CyclicBarrier) Await(ctx context.Context) error {
	b.lock.Lock()

	if b.round.isBroken {
		b.lock.Unlock()
		return ErrBrokenBarrier
	}

	b.round.count++

	count := b.round.count
	waitch := b.round.waitch
	brokech := b.round.brokech

	b.lock.Unlock()

	switch {
	case count > b.parties:
		panic(fmt.Errorf("parties is %d, access %d", count, b.parties))
	case count == b.parties:
		b.reset(true)
		return nil
	default:
		// count < b.parties
		select {
		case <-waitch:
			return nil
		case <-brokech:
			return ErrBrokenBarrier
		case <-ctx.Done():
			b.breakBarrier(true)
			return ctx.Err()
		}
	}
}

func (b *CyclicBarrier) breakBarrier(needLocked bool) {
	if needLocked {
		b.lock.Lock()
		defer b.lock.Unlock()
	}

	if !b.round.isBroken {
		close(b.round.brokech)
		b.round.isBroken = true
	}
}

func (b *CyclicBarrier) reset(safe bool) {
	b.lock.Lock()
	defer b.lock.Unlock()

	if safe {
		close(b.round.waitch)
	} else if b.round.count > 0 {
		b.breakBarrier(false)
	}

	b.round = &round{
		waitch:  make(chan struct{}),
		brokech: make(chan struct{}),
	}
}
