package general

import (
	"container/list"
	"fmt"
)

type LeakyBucket struct {
	TimeIntervalMs int64
	RequestCount   int

	Lookup map[string]*list.List
}

func NewLeakyBucket(RequestCount int, TimeIntervalMs int64) *LeakyBucket {
	return &LeakyBucket{
		RequestCount:   RequestCount,
		TimeIntervalMs: TimeIntervalMs,
		Lookup:         map[string]*list.List{},
	}
}

func (l *LeakyBucket) IsAllowed(key string, timestamp int64) bool {
	queue, found := l.Lookup[key]
	if !found {
		queue = list.New()
		l.Lookup[key] = queue
	}

	fmt.Printf("%v, %v, %v, \n", timestamp, queue, l.TimeIntervalMs)
	for queue.Len() > 0 && timestamp-queue.Front().Value.(int64) > l.TimeIntervalMs {
		fmt.Printf("%v, %v, %v, \n", timestamp, queue.Front().Value.(int64), l.TimeIntervalMs)
		queue.Remove(queue.Front())
	}

	fmt.Printf("%v, %v, \n", queue.Len(), l.RequestCount)
	if queue.Len() >= l.RequestCount {
		return false
	}

	queue.PushBack(timestamp)
	return true
}

func (l *LeakyBucket) Cleanup(olderThanTs int64) {
	removeKeys := []string{}
	for k, v := range l.Lookup {
		fmt.Printf("%v, %v, \n", v.Back().Value.(int64), olderThanTs)
		if v.Back().Value.(int64) < olderThanTs {
			removeKeys = append(removeKeys, k)
		}
	}

	fmt.Printf("%v, %v, \n", "", removeKeys)
	for _, k := range removeKeys {
		delete(l.Lookup, k)
	}
}
