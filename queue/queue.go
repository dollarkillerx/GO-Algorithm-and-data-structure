package queue

type MyQueue interface {
	Size() uint64
	First() interface{}
	End() interface{}
	IsEmpty() bool        // 是否为空
	EnQueue(interface{})  //  入列
	DeQueue() interface{} // 出列
	Clear()
}

type Queue struct {
	data []interface{} // 存储数据
}

func New() MyQueue {
	return &Queue{
		data: make([]interface{}, 0),
	}
}

func (m *Queue) Size() uint64 {
	return uint64(len(m.data))
}

func (m *Queue) First() interface{} {
	if len(m.data) > 0 {
		return m.data[0]
	}
	return nil
}

func (m *Queue) End() interface{} {
	if len(m.data) > 0 {
		return m.data[len(m.data)-1]
	}
	return nil
}

func (m *Queue) IsEmpty() bool {
	if len(m.data) == 0 {
		return true
	}
	return false
}

func (m *Queue) EnQueue(data interface{}) {
	m.data = append(m.data, data)
}

func (m *Queue) DeQueue() interface{} {
	if len(m.data) <= 0 {
		return nil
	}
	result := m.data[0]
	if len(m.data) > 1 {
		m.data = append(m.data[1:])
	} else {
		m.data = make([]interface{}, 0)
	}
	return result
}

func (m *Queue) Clear() {
	m.data = make([]interface{}, 0)
}
