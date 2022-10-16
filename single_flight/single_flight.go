package single_flight

import "sync"

type call struct {
	wg    sync.WaitGroup
	value interface{}
	err   error
}

type Group struct {
	m      sync.Mutex
	groups map[string]*call
}

func (g *Group) Do(name string, f func() (interface{}, error)) (interface{}, error) {
	g.m.Lock()
	if g.groups == nil {
		g.groups = make(map[string]*call)
	}
	c, ok := g.groups[name]
	if ok {
		g.m.Unlock()
		c.wg.Wait()
		return c.value, c.err
	}

	c = new(call)
	c.wg.Add(1)
	g.groups[name] = c
	g.m.Unlock()

	c.value, c.err = f()
	c.wg.Done()

	g.m.Lock()
	delete(g.groups, name)
	g.m.Unlock()
	
	return c.value, c.err
}
