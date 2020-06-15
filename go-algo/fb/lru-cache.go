package fb


import "container/list"
import "errors"


type node struct {
	key string
	value string
}

type NDCache struct {
	lookup map[string]*list.Element
	queue list.List
	capacity int
}


func (c *NDCache) get(key string) (string, error) {
	// validate // TODO

	//not present
	element, ok := c.lookup[key]
	if !ok {
		return "", errors.New("not found")
	}

	// move to front of queue
	c.queue.MoveToFront(element)

	// return from lookup
	return c.lookup[key].Value.(*node).value, nil // TODO
}

/////////////////////////////

func (c *NDCache) set(key string, value string) {
	// validate // TODO

	// if found then update
	_, err := c.get(key)
	if err == nil {
		c.update(key, value)
		return
	}

	// not found then check capacity
	// if full evict
	if c.queue.Len() >= c.capacity {
		c.evict()
	}

	// add to lookup and front of queue
	element := &node{
		key: key,
		value: value,
	}

	elem := c.queue.PushFront(element)
	c.lookup[key] = elem

}

func (c *NDCache) update(key string, value string) {
	// TODO
}

func (c *NDCache) evict() {
	// TODO

	elem := c.queue.Back()
	c.queue.Remove(elem)

	key := elem.Value.(*node).key
	delete(c.lookup, key)

}

/*
func main() {

	cache := NDCache{
		lookup: map[string]*list.Element{},
		queue: list.List{},
		capacity: 3,
	}

	cache.set("A", "W")

	value, err := cache.get("A")
	if err != nil {
		fmt.Printf("err get %v\n", err)
	} else {
		fmt.Printf("get %v\n", value)
	}

	fmt.Printf("cache %v\n", cache)

	cache.set("B", "X")
	cache.set("C", "Y")
	cache.set("D", "Z")

	value, err = cache.get("A")
	if err != nil {
		fmt.Printf("err get %v\n", err)
	} else {
		fmt.Printf("get %v\n", value)
	}

	value, err = cache.get("B")
	if err != nil {
		fmt.Printf("err get %v\n", err)
	} else {
		fmt.Printf("get %v\n", value)
	}

	fmt.Printf("cache %v\n", cache)



	cache.set("P", "Q")
	value, err = cache.get("B")
	if err != nil {
		fmt.Printf("err get %v\n", err)
	} else {
		fmt.Printf("get %v\n", value)
	}

	value, err = cache.get("C")
	if err != nil {
		fmt.Printf("err get %v\n", err)
	} else {
		fmt.Printf("get %v\n", value)
	}

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


