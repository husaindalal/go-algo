package cache

import (
	"container/list"
	"errors"
)

type Node struct {
	key   string
	value int //
}

type LRU struct {
	Lookup   map[string]*list.Element
	Queue    *list.List
	Capacity int
}

func New(capacity int) *LRU {
	return &LRU{
		Lookup:   make(map[string]*list.Element, capacity),
		Queue:    list.New(),
		Capacity: capacity,
	}
}

// Upsert
func (l *LRU) Add(key string, val int) {
	node := Node{key: key, value: val}
	// if full remove tail
	if l.Queue.Len() >= l.Capacity {
		l.RemoveOldest() // async
	}
	// if already in lookup, update the value and move to front
	if _, present := l.Lookup[node.key]; present {
		l.Update(node)
		return
	}

	ptr := l.Queue.PushFront(&list.Element{
		Value: node,
	})
	l.Lookup[node.key] = ptr

}

func (l *LRU) Get(key string) (int, error) {
	result, present := l.Lookup[key]
	if !present {
		return 0, errors.New("not present")
	}
	return result.Value.(*Node).value, nil
}

func (l *LRU) RemoveOldest() {
	elem := l.Queue.Back()
	delete(l.Lookup, elem.Value.(*list.Element).Value.(Node).key)

	l.Queue.Remove(elem)

}

func (l *LRU) Update(node Node) {
	// update value
	elem := l.Lookup[node.key]
	elem.Value.(*list.Element).Value = node
	// move to front
	l.Queue.MoveToFront(elem)
}
