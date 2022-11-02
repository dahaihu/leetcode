package sema

import (
	"container/list"
	"context"
	"sync"
)

type waiter struct {
	tickets int64
	ready   chan struct{}
}

type Weighted struct {
	mu   sync.Mutex
	list list.List
	cur  int64
	size int64
}

func NewWeighted(weighted int64) *Weighted {
	return &Weighted{size: weighted}
}

func (s *Weighted) Acquire(ctx context.Context, tickets int64) error {
	s.mu.Lock()
	if s.tickets() >= tickets && s.list.Len() == 0 {
		s.cur += tickets
		s.mu.Unlock()
		return nil
	}
	if tickets > s.size {
		s.mu.Unlock()
		<-ctx.Done()
		return ctx.Err()
	}
	ready := make(chan struct{})
	w := waiter{tickets: tickets, ready: ready}
	elem := s.list.PushBack(w)
	s.mu.Unlock()

	select {
	case <-ready:
		return nil
	case <-ctx.Done():
		err := ctx.Err()
		s.mu.Lock()
		select {
		case <-ready:
			err = nil
		default:
			isFront := s.list.Front() == elem
			s.list.Remove(elem)
			if isFront && s.tickets() > 0 {
				s.notify()
			}
		}
		s.mu.Unlock()
		return err
	}
}

func (s *Weighted) Release(tickets int64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.cur -= tickets
	if s.cur < 0 {
		panic("semaphore: released more than held")
	}
	s.notify()
}

func (s *Weighted) notify() {
	for {
		front := s.list.Front()
		if front == nil {
			break
		}
		waiter := front.Value.(waiter)
		if s.size-s.cur < waiter.tickets {
			break
		}
		s.list.Remove(front)
		s.cur += waiter.tickets
		close(waiter.ready)
	}
}

func (s *Weighted) tickets() int64 {
	return s.size - s.cur
}
