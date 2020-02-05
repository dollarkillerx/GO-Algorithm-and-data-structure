package link_queue

import (
	"errors"
)

type LinkQueueInterface interface {
	Size() uint64
	IsEmpty() bool                 // 是否为空
	EnQueue(interface{})           //  入列
	DeQueue() (interface{}, error) // 出列
	Clear()
}

type linkQueueNode struct {
	data  interface{}
	pNext *linkQueueNode
}

type LinkQueue struct {
	header *linkQueueNode
	footer *linkQueueNode
	len    uint64
}

func New() LinkQueueInterface {
	return &LinkQueue{
		header: nil,
		footer: nil,
		len:    0,
	}
}

func (l *LinkQueue) Size() uint64 {
	return l.len
}

func (l *LinkQueue) IsEmpty() bool {
	if l.header == nil {
		return true
	}
	return false
}

func (l *LinkQueue) EnQueue(data interface{}) {
	l.len++
	node := &linkQueueNode{
		data:  data,
		pNext: nil,
	}
	if l.header != nil {
		l.footer.pNext = node
		l.footer = node
	} else {
		l.header = node
		l.footer = node
	}

}

func (l *LinkQueue) DeQueue() (interface{}, error) {
	if l.IsEmpty() {
		return nil, errors.New("is empty")
	}
	defer func() {
		l.len--
	}()
	if l.len == 1 {
		result := l.header
		l.header = nil
		l.footer = nil
		return result.data, nil
	} else {
		result := l.header
		l.header = l.header.pNext
		return result.data, nil
	}
}

func (l *LinkQueue) Clear() {
	l.header = nil
	l.footer = nil
	l.len = 0
}
