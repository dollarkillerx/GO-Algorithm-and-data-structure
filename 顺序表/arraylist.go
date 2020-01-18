/**
*@program: GO-Algorithm-and-data-structure
*@description: https://github.com/dollarkillerx
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-01-16 12:59
 */
package array

import (
	"encoding/json"
	"errors"
)

type List interface {
	Size() int                             // 返回list中数据的size
	Get(index int) (interface{}, error)    // 获取下标为index 的数据
	Set(index int, data interface{}) error // set
	Append(data interface{})               // 向list最末尾添加数据
	Clear()                                // 清空
	Delete(index int) error                // delete
	String() string                        // to string
}

type ArrayList struct {
	dataStore []interface{}
	size      int // list size
}

func NewArrayList() *ArrayList {
	return &ArrayList{
		dataStore: make([]interface{}, 10),
		size:      0,
	}
}

func (a *ArrayList) Size() int {
	return a.size
}

func (a *ArrayList) Get(index int) (interface{}, error) {
	if index >= a.size || index < 0 {
		return nil, errors.New("Subscript out of bounds")
	}
	return a.dataStore[index], nil
}

func (a *ArrayList) Set(index int, data interface{}) error {
	if index >= len(a.dataStore) || index < 0 {
		return errors.New("Subscript out of bounds")
	}
	a.dataStore[index] = data
	a.size++
	return nil
}

func (a *ArrayList) Append(data interface{}) {
	if a.size >= len(a.dataStore) {
		// 扩容
		ne := make([]interface{}, len(a.dataStore)+len(a.dataStore)/2)
		for k, v := range a.dataStore {
			if v == nil {
				break
			}
			ne[k] = v
		}
		a.dataStore = ne

	}
	// 插入
	a.dataStore[a.size] = data
	a.size++
}

func (a *ArrayList) Clear() {
	a.dataStore = make([]interface{}, 0, 0)
	a.size = 0
}

func (a *ArrayList) Delete(index int) error {
	if index >= a.size || index < 0 {
		return errors.New("Subscript out of bounds")
	}
	a.size--
	if index != a.size-1 {
		a.dataStore = append(a.dataStore[:index], a.dataStore[index+1:]...)
	} else {
		a.dataStore = append(a.dataStore[:index])
	}
	return nil
}

func (a *ArrayList) String() string {
	marshal, _ := json.Marshal(a.dataStore)
	return string(marshal)
}
