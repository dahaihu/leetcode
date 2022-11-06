package lfu

import (
	"bytes"
	"container/list"
	"fmt"
)

type valueEntry struct {
	key, value int
	freqNode   *list.Element
}

func newValueEntry(key, value int, freqNode *list.Element) *valueEntry {
	return &valueEntry{key: key, value: value, freqNode: freqNode}
}

type frequencyEntry struct {
	frequency int
	valueList *list.List
}

func newfrequencyEntry(freq int) *frequencyEntry {
	return &frequencyEntry{
		frequency: freq,
		valueList: list.New(),
	}
}

type LFUCache struct {
	capacity   int
	freqList   *list.List
	valueNodes map[int]*list.Element
}

func LFUConstructor(cap int) *LFUCache {
	if cap <= 0 {
		panic("invalid lru capacity")
	}
	return &LFUCache{
		capacity:   cap,
		freqList:   list.New(),
		valueNodes: make(map[int]*list.Element),
	}
}

func (lfu *LFUCache) String() string {
	var buffer bytes.Buffer
	freqNode := lfu.freqList.Front()
	var outerdep, innerdep int
	for freqNode != nil {
		outerdep++
		innerdep = 0
		freqEntry := freqNode.Value.(*frequencyEntry)
		buffer.WriteString(fmt.Sprintf("freq: [%d] ", freqEntry.frequency))
		valueNode := freqEntry.valueList.Front()
		for valueNode != nil {
			innerdep++
			valueEntry := valueNode.Value.(*valueEntry)
			buffer.WriteString(fmt.Sprintf("->(%3d, %3d)",
				valueEntry.key, valueEntry.value))
			valueNode = valueNode.Next()
		}
		buffer.WriteString("\n")
		freqNode = freqNode.Next()
	}
	return buffer.String()
}

func (lfu *LFUCache) evict() {
	freqNode := lfu.freqList.Front()
	freqEntry := freqNode.Value.(*frequencyEntry)

	removedNode := freqEntry.valueList.Back()
	removedEntry := removedNode.Value.(*valueEntry)

	freqEntry.valueList.Remove(removedNode)
	delete(lfu.valueNodes, removedEntry.key)

	if freqEntry.valueList.Len() == 0 {
		lfu.freqList.Remove(freqNode)
	}
}

func (lfu *LFUCache) increment(valueNode *list.Element) {
	entry := valueNode.Value.(*valueEntry)

	freqNode := entry.freqNode
	freqEntry := freqNode.Value.(*frequencyEntry)

	// get next freq node and freq entry
	nextFreqNode := freqNode.Next()
	var nextFreqEntry *frequencyEntry
	if nextFreqNode == nil || nextFreqNode.Value.(*frequencyEntry).frequency != freqEntry.frequency+1 {
		nextFreqEntry = &frequencyEntry{frequency: freqEntry.frequency + 1, valueList: list.New()}
		nextFreqNode = lfu.freqList.InsertAfter(nextFreqEntry, freqNode)
	} else {
		nextFreqEntry = nextFreqNode.Value.(*frequencyEntry)
	}

	freqEntry.valueList.Remove(valueNode)
	entry.freqNode = nextFreqNode

	newValueNode := nextFreqEntry.valueList.PushBack(entry)
	lfu.valueNodes[entry.key] = newValueNode

	if freqEntry.valueList.Len() == 0 {
		lfu.freqList.Remove(freqNode)
	}
}

func (lfu *LFUCache) Put(key, value int) {
	if valueNode, ok := lfu.valueNodes[key]; !ok {
		if len(lfu.valueNodes) == lfu.capacity {
			lfu.evict()
		}
		freqNodeHead := lfu.freqList.Front()
		if freqNodeHead == nil {
			newFreqEntry := newfrequencyEntry(1)
			newFreqNode := lfu.freqList.PushFront(newFreqEntry)

			valueEntry := newValueEntry(key, value, newFreqNode)
			valueNode := newFreqEntry.valueList.PushFront(valueEntry)
			lfu.valueNodes[key] = valueNode
			return
		}
		freqNodeHeadEntry := freqNodeHead.Value.(*frequencyEntry)
		if freqNodeHeadEntry.frequency != 1 {
			freqNodeHeadEntry = newfrequencyEntry(1)
			freqNodeHead = lfu.freqList.PushFront(freqNodeHeadEntry)
		}
		valueEntry := newValueEntry(key, value, freqNodeHead)
		nodeEntry := freqNodeHeadEntry.valueList.PushFront(valueEntry)
		lfu.valueNodes[key] = nodeEntry
	} else {
		lfu.increment(valueNode)
		valueEntry := valueNode.Value.(*valueEntry)
		valueEntry.value = value
	}
}

func (lfu *LFUCache) Get(key int) int {
	if valueNode, ok := lfu.valueNodes[key]; ok {
		lfu.increment(valueNode)
		entry := valueNode.Value.(*valueEntry)
		return entry.value
	}
	return -1
}
