package queue

import (
	"fmt"
	"testing"
)

func TestQueue_Clear(t *testing.T) {
	queue := New()

	fmt.Println(queue.IsEmpty())
	fmt.Println(queue.Size())

	queue.EnQueue("this is ok")
	queue.EnQueue("ok ")

	fmt.Println(queue.First())
	fmt.Println(queue.End())

	fmt.Println(queue.DeQueue())

	queue.EnQueue("2020")

	fmt.Println(queue.Size())
	fmt.Println(queue.IsEmpty())

	queue.Clear()

	fmt.Println(queue.Size())
	fmt.Println(queue.IsEmpty())
}

func TestTwo(t *testing.T) {
	a := 0
	a++
	fmt.Println(a)
}
