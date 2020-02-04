package queue

type MyQueue interface {
	Size() uint64
	First() interface{}
	End() interface{}
	IsEmpty() bool // 是否为空
	EnQueue(interface{}) //  入列
	DeQueue() interface{} // 出列
	Clear()
}

type Queue struct {
	data []interface{} // 存储数据
	len uint64 // 队列存储
}

func New() MyQueue  {
	return &Queue{
		data: make([]interface{},0),
		len:0,
	}
}

func (m *Queue) Size() uint64 {
	return m.len
}

func (m *Queue) First() interface{} {
	if m.len > 0 {
		return m.data[0]
	}
	return nil
}

func (m *Queue) End() interface{} {
	if m.len > 0 {
		return m.data[m.len-1]
	}
	return nil
}

func (m *Queue) IsEmpty() bool {
	if m.len == 0 {
		return true
	}
	return false
}

func (m *Queue) EnQueue(data interface{})  {
	m.data = append(m.data,data)
	m.len ++
}

func (m *Queue) DeQueue() interface{} {
	if m.len <= 0 {
		return nil
	}
	result := m.data[0]
	if m.len > 1 {
		m.data = append(m.data[1:])
	}else {
		m.data = make([]interface{},0)
	}
	return result
}

func (m *Queue) Clear()  {
	m.len = 0
	m.data = make([]interface{},0)
}