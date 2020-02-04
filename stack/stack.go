package stack

import (
	"errors"
)

type StackArrayInterface interface {
	Push(data interface{}) error
	Pop() (interface{}, error)
	Clear()
	Size() uint64
	IsFull() bool  // 是否满了
	IsEmpty() bool // 是否为空
}

type Stack struct {
	array []interface{}
	len   uint64
	cap   uint64
}

func New(cap uint64) StackArrayInterface {
	//func New(cap uint64) *Stack {
	return &Stack{
		array: make([]interface{}, 0),
		len:   0,
		cap:   cap,
	}
}

func (s *Stack) Push(data interface{}) error {
	if s.len > s.cap {
		return errors.New("stack full")
	}
	s.len++
	s.array = append(s.array, data)
	return nil
}

func (s *Stack) Pop() (interface{}, error) {
	if s.len <= 0 {
		return nil, errors.New("not data")
	}
	defer func() {
		s.len--
		s.array = append(s.array[:s.len])
	}()
	return s.array[s.len-1], nil
}

func (s *Stack) Clear() {
	s.len = 0
	s.array = make([]interface{}, s.cap)
}

func (s *Stack) Size() uint64 {
	return s.len
}

func (s *Stack) IsFull() bool {
	if s.cap == s.len {
		return true
	}
	return false
}

func (s *Stack) IsEmpty() bool {
	if s.len == 0 {
		return true
	}
	return false
}

func (s *Stack) GetAR() []interface{} {
	return s.array
}
