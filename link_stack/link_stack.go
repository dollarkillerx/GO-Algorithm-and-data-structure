package link_stack

import "errors"

type LinkStackInterface interface {
	Push(data interface{})
	Pop() (interface{}, error)
	Clear()
	Size() uint64
	IsEmpty() bool
}

type LinkStack struct {
	data  interface{}
	pNext *LinkStack
	len   uint64
}

func New() LinkStackInterface {
	return &LinkStack{
		data:  nil,
		pNext: nil,
	}
}

func (l *LinkStack) Push(data interface{}) {
	newStack := &LinkStack{
		data:  data,
		pNext: l.pNext,
	}

	l.pNext = newStack
	l.len++
}

func (l *LinkStack) Pop() (interface{}, error) {
	if l.IsEmpty() {
		return nil, errors.New("is empty")
	}
	result := l.pNext.data
	l.pNext = l.pNext.pNext
	return result, nil
}

func (l *LinkStack) Clear() {
	l.pNext = nil
	l.data = nil
	l.len = 0
}

func (l *LinkStack) Size() uint64 {
	return l.len
}

func (l *LinkStack) IsEmpty() bool {
	if l.pNext == nil {
		return true
	}
	return false
}
