package scheduler

import "sync/atomic"

type IJob interface {
	// Do() real do function
	Do()
	// ExecuteTime() return job's execute time second timestamp
	ExecuteTime() int64
}

// JobWrapper, wraps the real job
type JobWrapper struct {
	job     IJob
	index   atomic.Int64
	removed atomic.Bool
}

// implements IElement Priority() method
func (j *JobWrapper) Priority() int64 {
	return j.job.ExecuteTime()
}

// implements IElement SetIndex() method
func (j *JobWrapper) SetIndex(idx int) {
	j.index.Store(int64(idx))
}

// implements IElement Index() method
func (j *JobWrapper) Index() int {
	idx := j.index.Load()
	return int(idx)
}

// Remove, remove job from priority queue
func (j *JobWrapper) Remove() {
	j.removed.Store(true)
}

// Removed, check job is removed from priority queue
func (j *JobWrapper) Removed() bool {
	return j.removed.Load()
}

// NewJobWrapper, create a job from IJob
func NewJobWrapper(job IJob) *JobWrapper {
	return &JobWrapper{job: job}
}
