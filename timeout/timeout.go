package timeout

import (
	"errors"
	"time"
)

var ErrTimeout = errors.New("timeout")

func TimeoutFunc(duration time.Duration, f func() interface{}) (interface{}, error) {
	c := make(chan interface{}, 1)

	go func() {
		c <- f()
		close(c)
	}()

	select {
	case <-time.After(duration):
		return nil, ErrTimeout
	case val := <-c:
		return val, nil
	}
}
