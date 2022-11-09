package cycle_print

var token = struct{}{}

type Cycle struct {
	chls []chan struct{}
	f    func(workerIdx int, times int) bool // times signigi
	done chan struct{}
	cur  int // initiate value is zero, every time a worker has done a work, it increment one
}

func New(workerNums int, f func(workerIdx int, times int) bool) *Cycle {
	if workerNums <= 0 {
		panic("invalid workerNums in NewCycle")
	}
	chs := make([]chan struct{}, workerNums)
	for i := 0; i < workerNums; i++ {
		chs[i] = make(chan struct{})
	}
	cycle := &Cycle{
		chls: chs,
		f:    f,
		done: make(chan struct{}),
	}
	return cycle
}

func (c *Cycle) Start() {
	for i := 0; i < len(c.chls); i++ {
		go c.cycle(i)
	}
	c.notify(0)
}

func (c *Cycle) Done() <-chan struct{} {
	return c.done
}

func (c *Cycle) wait(i int) (ok bool) {
	_, ok = <-c.chls[i]
	return ok
}

func (c *Cycle) notify(i int) {
	if i == len(c.chls) {
		i = 0
	}
	c.chls[i] <- token
}

func (c *Cycle) end() {
	for _, ch := range c.chls {
		close(ch)
	}
	close(c.done)
}

func (c *Cycle) cycle(idx int) {
	for {
		ok := c.wait(idx)
		if !ok {
			break
		}
		if shouldContinue := c.f(idx, c.cur); !shouldContinue {
			c.end()
			break
		}
		c.cur++
		c.notify(idx + 1)
	}
}
