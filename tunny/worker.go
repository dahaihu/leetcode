package tunny

type Worker interface {
	Process(interface{}) interface{}
	Interrupt()
	BlockUntilReady()
	Terminate()
}

type workRequest struct {
	payloadChan chan interface{}
	retChan     chan interface{}
	interrupt   func()
}

type workerWrapper struct {
	requestChan   chan *workRequest
	worker        Worker
	interruptChan chan struct{}
	closeChan     chan struct{}
	closedChan    chan struct{}
}

func newWorkerWrapper(reqChan chan *workRequest, worker Worker) *workerWrapper {
	w := &workerWrapper{
		requestChan:   reqChan,
		worker:        worker,
		interruptChan: make(chan struct{}),
		closeChan:     make(chan struct{}),
		closedChan:    make(chan struct{}),
	}
	go w.work()
	return w
}

func (w *workerWrapper) interrupt() {
	w.worker.Interrupt()
	close(w.interruptChan)
}

func (w *workerWrapper) stop() {
	w.worker.Terminate()
	close(w.closeChan)
}

func (w *workerWrapper) join() {
	<-w.closedChan
}

func (w *workerWrapper) work() {
	payloadChan, retChan := make(chan interface{}), make(chan interface{})
	defer func() {
		w.worker.Terminate()
		close(retChan)
		close(w.closedChan)
	}()
	for {
		w.worker.BlockUntilReady()
		select {
		case w.requestChan <- &workRequest{
			payloadChan: payloadChan,
			retChan:     retChan,
			interrupt:   w.interrupt,
		}:
			select {
			case payload := <-payloadChan:
				ret := w.worker.Process(payload)
				select {
				case retChan <- ret:
				case <-w.interruptChan:
					w.interruptChan = make(chan struct{})
				}
			case <-w.interruptChan:
				w.interruptChan = make(chan struct{})
			}
		case <-w.closeChan:
			return
		}
	}
}
