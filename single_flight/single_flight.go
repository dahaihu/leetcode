package single_flight

import "sync"

type call struct {
	val interface{}
	err error
	wg  sync.WaitGroup
}

type Group struct {
	calls map[string]*call
	m     sync.Mutex
}

func (g *Group) Do(key string, f func() (interface{}, error)) (interface{}, error) {
	g.m.Lock()

	if g.calls == nil {
		g.calls = make(map[string]*call)
	}
	c, ok := g.calls[key]
	if ok {
		g.m.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}
	c = new(call)
	c.wg.Add(1)
	g.calls[key] = c
	g.m.Unlock()

	c.val, c.err = f()
	c.wg.Done()

	g.m.Lock()
	delete(g.calls, key)
	g.m.Unlock()

	return c.val, c.err
}
