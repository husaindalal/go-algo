package fb

import (
	"container/list"
)

type node struct {
	key   int
	value string
}

type NextCache struct {
	lookup   map[int]*list.Element
	queue    list.List // nodes
	capacity int
	// mutex
	//mutex sync.RWMutex
}

func NewCache(capacity int) *NextCache {
	return &NextCache{
		lookup:   map[int]*list.Element{},
		queue:    list.List{},
		capacity: capacity,
		//mutex:    sync.RWMutex{},
	}
}

func (n *NextCache) LRUPut(key int, value string) string {

	//n.mutex.Lock()
	//defer n.mutex.Unlock()

	// already present
	data := n.LRUGet(key)
	if data != "" {
		// TODO update
		return data
	}

	// not present
	// if queue full, remove tail
	if len(n.lookup) == n.capacity {
		n.removeTail()
	}

	element := &list.Element{
		Value: node{
			key:   key,
			value: value,
		},
	}
	// add element to head of queue
	n.queue.PushFront(element)

	// add element to lookup
	n.lookup[key] = element

	return value
}

func (n *NextCache) LRUGet(key int) string {

	//n.mutex.RLock()
	//defer n.mutex.RUnlock()

	// check in lookup
	element, ok := n.lookup[key]

	// if not present return ""
	if !ok {
		return "" // TODO throw err
	}

	// if present move to front of queue
	n.queue.MoveToFront(element)

	// return value
	return element.Value.(node).value
}

func (n *NextCache) removeTail() {
	element := n.queue.Back()
	n.queue.Remove(element)
	delete(n.lookup, element.Value.(*list.Element).Value.(node).key)
}

/*


func main() {

	cache := fb.NewCache(3)

	fmt.Printf("LRUGet %v \n", cache.LRUGet(1))

	fmt.Printf("LRUPut %v \n", cache.LRUPut(1, "A"))
	fmt.Printf("LRUGet %v \n", cache.LRUGet(1))

	fmt.Printf("LRUPut %v \n", cache.LRUPut(2, "B"))
	fmt.Printf("LRUPut %v \n", cache.LRUPut(3, "C"))
	fmt.Printf("LRUPut %v \n", cache.LRUPut(4, "D"))

	fmt.Printf("LRUGet %v \n", cache.LRUGet(1))
	fmt.Printf("LRUGet %v \n", cache.LRUGet(2))
	//fmt.Printf("LRUGet %v \n", cache.LRUGet(1))


}

*/
