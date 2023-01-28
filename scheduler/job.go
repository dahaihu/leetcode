package scheduler

import "sync/atomic"

type IJob interface {
	// Do() real job function
	Do()
	// ExecuteTime() return job's execute time second timestamp
	ExecuteTime() int64
}

// JobWrapper, wraps the real job
type JobWrapper struct {
	job     IJob
	index   int64
	removed int64
}

// Priority implements IElement Priority() method
func (j *JobWrapper) Priority() int64 {
	return j.job.ExecuteTime()
}

// SetIndex implements IElement SetIndex() method
func (j *JobWrapper) SetIndex(idx int) {
	atomic.StoreInt64(&j.index, int64(idx))
}

// Index implements IElement Index() method
func (j *JobWrapper) Index() int {
	return int(atomic.LoadInt64(&j.index))
}

// Remove remove job from priority queue
func (j *JobWrapper) Remove() {
	atomic.StoreInt64(&j.removed, 1)
}

// Removed, check job is removed from priority queue
func (j *JobWrapper) Removed() bool {
	return atomic.LoadInt64(&j.removed) == 1
}

// NewJobWrapper, create a job from IJob
func NewJobWrapper(job IJob) *JobWrapper {
	return &JobWrapper{job: job}
}
