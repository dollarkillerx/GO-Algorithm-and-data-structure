package circle_queue

import "errors"

//
type CircleQueueInterface interface {
	Size() uint64
	IsEmpty() bool                 // 是否为空
	IsFull() bool                  // 是否满了
	EnQueue(interface{}) error     //  入列
	DeQueue() (interface{}, error) // 出列
	Clear()
}

type CircleQueue struct {
	data   []interface{}
	header uint64
	tail   uint64
	len    uint64
	cap    uint64
}

func New(cap uint64) CircleQueueInterface {
	return &CircleQueue{
		data:   make([]interface{}, cap),
		header: 0,
		tail:   0,
		len:    0,
		cap:    cap,
	}
}

func (c *CircleQueue) Size() uint64 {
	return c.len
}

func (c *CircleQueue) IsEmpty() bool {
	if c.len <= 0 {
		return true
	}
	return false
}

func (c *CircleQueue) IsFull() bool {
	if c.len >= c.cap {
		return true
	}
	return false
}

func (c *CircleQueue) EnQueue(data interface{}) error {
	if c.IsFull() {
		return errors.New("is full")
	}
	c.data[c.tail] = data
	if c.tail == c.cap-1 {
		c.tail = 0
	} else {
		c.tail++
	}
	c.len++
	return nil
}

func (c *CircleQueue) DeQueue() (interface{}, error) {
	if c.IsEmpty() {
		return nil, errors.New("is empty")
	}
	result := c.data[c.header]

	if c.header == c.cap-1 {
		c.header = 0
	} else {
		c.header++
	}
	c.len--
	return result, nil
}

func (c *CircleQueue) Clear() {
	c.len = 0
	c.header = 0
	c.tail = 0
	for k := range c.data {
		c.data[k] = 0
	}
}
