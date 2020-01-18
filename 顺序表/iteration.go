/**
*@program: iteration 迭代器
*@description: https://github.com/dollarkillerx
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-01-18 10:00
 */
package array

import (
	"errors"
)

type Iteration interface {
	IsNext() bool
	Next() (interface{}, error)
	Reset()
}

type Iterable interface {
	Iteration() Iteration
}

type ArrayListIteration struct {
	data  *ArrayList
	index int
}

func (a *ArrayList) Iteration() Iteration {
	return &ArrayListIteration{
		data:  a,
		index: 0,
	}
}

func (a *ArrayListIteration) IsNext() bool {
	if a.index >= a.data.size {
		return false
	}
	return true
}

func (a *ArrayListIteration) Next() (interface{}, error) {
	if a.index >= a.data.size {
		return nil, errors.New("下标越界")
	}
	defer func() {
		a.index++
	}()
	return a.data.Get(a.index)
}
func (a *ArrayListIteration) Reset() {
	a.index = 0
}
