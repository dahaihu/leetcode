package batch

import "time"

type Timer struct {
	time.Timer
}

func (t *Timer) Reset(duration time.Duration) {
	t.Timer.Reset(duration)
	select {
	case <-t.C:
	default:
	}
}
