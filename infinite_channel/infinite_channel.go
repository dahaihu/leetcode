package infinite_channel

import (
	"container/list"
)

type InfiniteChan struct {
	list *list.List
	In   chan interface{}
	Out  chan interface{}
}

func New() *InfiniteChan {
	ic := &InfiniteChan{
		list: list.New(),
		In:   make(chan interface{}),
		Out:  make(chan interface{}),
	}
	go ic.run()
	return ic
}

func (i *InfiniteChan) run() {
	for {
		if i.list.Len() == 0 {
			ele, ok := <-i.In
			if !ok {
				goto end
			}
			i.list.PushBack(ele)
		}

		head := i.list.Front()
		select {
		case i.Out <- head.Value:
			i.list.Remove(head)
		case ele, ok := <-i.In:
			if !ok {
				goto flush
			}
			i.list.PushBack(ele)
		}
	}

flush:
	for i.list.Len() != 0 {
		front := i.list.Front()
		i.Out <- front.Value
		i.list.Remove(front)
	}
end:
	close(i.Out)
}

func (i *InfiniteChan) Close() {
	close(i.In)
}
