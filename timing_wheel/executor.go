package timing_wheel

type Executor struct {
	capacity int
	name     string
	sema     chan struct{}
}

func NewExecutor(name string, capacity int) *Executor {
	var sema chan struct{}
	if capacity > 0 {
		sema = make(chan struct{}, capacity)
	}
	return &Executor{
		capacity: capacity,
		name:     name,
		sema:     sema,
	}
}

func (e *Executor) Execute(task *TimerTask) {
	if e.sema != nil {
		l.Infof("executor[%s] waiting for sema for task[%s]", e.name, task.Key())
		e.sema <- struct{}{}
		l.Infof("executor[%s] get sema for task[%s]", e.name, task.Key())
		go func() {
			defer func() {
				<-e.sema
			}()

			e.do(task)
		}()
	} else {
		go e.do(task)
	}
}

func (e *Executor) do(task *TimerTask) {
	l.Infof("executor[%s] start do task[%s]", e.name)
	defer func() {
		if err := recover(); err != nil {
			l.Errorf("executor[%s] do task[%s] panic: %+v", e.name, task.Key(), err)
		}
		l.Infof("executor[%s] finish do task[%s]", e.name, task.Key())
	}()

	task.Do()
}
