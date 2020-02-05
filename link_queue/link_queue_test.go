package link_queue

import (
	"fmt"
	"strconv"
	"testing"
)

func TestLinkQueue_Clear(t *testing.T) {
	queue := New()

	fmt.Println(queue.IsEmpty())
	fmt.Println(queue.Size())

	for i := 0; i < 100; i++ {
		queue.EnQueue("is ok" + strconv.Itoa(i))
	}

	fmt.Println(queue.IsEmpty())
	fmt.Println(queue.Size())

	for i := 0; i < 100; i++ {
		deQueue, e := queue.DeQueue()
		if e != nil {
			fmt.Println(e)
			continue
		}
		fmt.Println(deQueue)
	}

	fmt.Println(queue.IsEmpty())
	fmt.Println(queue.Size())

	for i := 0; i < 30; i++ {
		queue.EnQueue("is ok pro: " + strconv.Itoa(i))
	}

	fmt.Println(queue.IsEmpty())
	fmt.Println(queue.Size())

	queue.Clear()

	fmt.Println(queue.IsEmpty())
	fmt.Println(queue.Size())
}
