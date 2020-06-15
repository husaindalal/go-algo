package queue

import "errors"

//import "container/list"

type Item int

type CircularQueue struct {
	queue    []*Item
	capacity int
	head     int
	tail     int
}

func (c *CircularQueue) New(capacity int) {
	c.queue = make([]*Item, capacity)
	c.capacity = capacity
	c.head = 0
}

func (c *CircularQueue) IsFull() bool {
	if c.head == (c.tail+1)%c.capacity {
		return true
	}
	return false
}

func (c *CircularQueue) IsEmpty() bool {
	if c.head == c.tail {
		return true
	}
	return false
}

func (c *CircularQueue) Enqueue(item Item) error {
	if c.IsFull() {
		return errors.New("queue full")
	}
	c.queue[c.tail] = &item
	c.tail = (c.tail + 1) % c.capacity

	return nil
}

func (c *CircularQueue) Dequeue() (*Item, error) {
	if c.IsEmpty() {
		return nil, errors.New("queue empty")
	}
	item := c.queue[c.head]
	c.head = (c.head + 1) % c.capacity
	return item, nil
}
