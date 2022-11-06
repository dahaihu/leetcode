package retry

import (
	"context"
	"errors"
	"time"
)

func ConstantBackoff(n int, backoff time.Duration) []time.Duration {
	durations := make([]time.Duration, n)
	for i := 0; i < n; i++ {
		durations[i] = backoff
	}
	return durations
}

func ExponentialBackoff(n int, initBackoff time.Duration) []time.Duration {
	durations := make([]time.Duration, n)
	durations[0] = initBackoff
	for i := 1; i < n; i++ {
		durations[i] = durations[i-1] * durations[i-1]
	}
	return durations
}

func LimitExponentialBackoff(n int, initBackoff time.Duration, limitBackoff time.Duration) []time.Duration {
	durations := make([]time.Duration, n)
	durations[0] = initBackoff
	for i := 1; i < n; i++ {
		cur := durations[i-1] * durations[i-1]
		if cur <= limitBackoff {
			durations[i] = cur
		} else {
			for j := i; j < n; j++ {
				durations[j] = limitBackoff
			}
			break
		}
	}
	return durations
}

type ErrorAction int

const (
	Succeed ErrorAction = iota
	Fail
	Retry
)

type IErrorClassfier interface {
	Classfy(error) ErrorAction
}

type DefaultClassifier struct{}

func (c DefaultClassifier) Classfy(err error) ErrorAction {
	if err == nil {
		return Succeed
	}
	return Retry
}

type WhitelistClassfier []error

func (c WhitelistClassfier) Classfy(err error) ErrorAction {
	if err == nil {
		return Succeed
	}
	for _, mark := range c {
		if errors.Is(mark, err) {
			return Retry
		}
	}
	return Fail
}

type BlacklistClassfier []error

func (c BlacklistClassfier) Classfy(err error) ErrorAction {
	if err == nil {
		return Succeed
	}
	for _, mark := range c {
		if errors.Is(err, mark) {
			return Fail
		}
	}
	return Retry
}

type Retrier struct {
	backoff   []time.Duration
	classfier IErrorClassfier
}

func New(backoffs []time.Duration, classfier IErrorClassfier) *Retrier {
	if len(backoffs) == 0 {
		panic("invalid backoffs, no need retry")
	}
	if classfier == nil {
		classfier = DefaultClassifier{}
	}
	return &Retrier{
		backoff:   backoffs,
		classfier: classfier,
	}
}

func (r *Retrier) Run(work func(context.Context) error) error {
	return r.RunCtx(context.Background(), work)
}

func (r *Retrier) RunCtx(ctx context.Context, work func(context.Context) error) error {
	retries := 0
	for {
		ret := work(ctx)
		switch err := r.classfier.Classfy(ret); err {
		case Succeed, Fail:
			return ret
		case Retry:
			if retries >= len(r.backoff) {
				return ret
			}
			d := r.backoff[retries]
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.NewTimer(d).C:
			}
		}
	}
}
