package tunny

type Worker interface {
	Process(interface{}) interface{}
	Interrupt()
	BlockUntilReady()
	Terminate()
}

type workRequest struct {
	PayloadChan chan interface{}
	RetChan     chan interface{}
	Interrupt   func()
}

type workerWrapper struct {
	reqChan       chan<- workRequest
	worker        Worker
	interruptChan chan struct{}
	closeChan     chan struct{}
	closedChan    chan struct{}
}

func newWorkerWrapper(reqChan chan<- workRequest, worker Worker) *workerWrapper {
	w := &workerWrapper{
		reqChan:       reqChan,
		worker:        worker,
		interruptChan: make(chan struct{}),
		closeChan:     make(chan struct{}),
		closedChan:    make(chan struct{}),
	}
	go w.run()
	return w
}

func (w *workerWrapper) run() {
	payloadChan, retChan := make(chan interface{}), make(chan interface{})
	defer func() {
		w.worker.Terminate()
		close(retChan)
		close(w.closedChan)
	}()
	for {
		w.worker.BlockUntilReady()
		select {
		case w.reqChan <- workRequest{
			PayloadChan: payloadChan,
			RetChan:     retChan,
			Interrupt:   w.interrupt,
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

func (w *workerWrapper) interrupt() {
	w.worker.Interrupt()
	close(w.interruptChan)
}

func (w *workerWrapper) stop() {
	close(w.closeChan)
}

func (w *workerWrapper) join() {
	<-w.closedChan
}
