package single_flight

import "sync"

type call struct {
	val interface{}
	err error
	wg  sync.WaitGroup
}

type Group struct {
	callers map[string]*call
	m       sync.Mutex
}

func (g *Group) Do(name string, f func() (interface{}, error)) (interface{}, error) {
	g.m.Lock()
	if g.callers == nil {
		g.callers = make(map[string]*call)
	}
	c, ok := g.callers[name]
	if ok {
		g.m.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}
	c = new(call)
	c.wg.Add(1)
	g.callers[name] = c
	g.m.Unlock()

	c.val, c.err = f()
	c.wg.Done()

	g.m.Lock()
	delete(g.callers, name)
	g.m.Unlock()

	return c.val, c.err
}
