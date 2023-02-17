package batch

type sema chan struct{}

func newSema(concurrency int) sema {
	if concurrency > 0 {
		return make(chan struct{}, concurrency)
	}
	return nil
}

func (s sema) acquire() {
	if s != nil {
		s <- struct{}{}
	}
}

func (s sema) release() {
	if s != nil {
		<-s
	}
}
