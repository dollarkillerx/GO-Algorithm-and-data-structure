package circle_queue

import (
	"log"
	"strconv"
	"testing"
)

func TestCircleQueue_Clear(t *testing.T) {
	queue := New(100)
	for i := 0; i < 60; i++ {
		err := queue.EnQueue("is ok" + strconv.Itoa(i))
		if err != nil {
			log.Println(err)
			continue
		}
	}

	for i := 0; i < 30; i++ {
		deQueue, e := queue.DeQueue()
		if e != nil {
			log.Println(e)
			continue
		}
		log.Println(deQueue)
	}

	for i := 0; i < 100; i++ {
		err := queue.EnQueue("Go go go" + strconv.Itoa(i))
		if err != nil {
			log.Println(err)
			continue
		}
	}

	log.Println()
	log.Println()
	log.Println(queue.IsEmpty())
	log.Println(queue.IsFull())
	log.Println(queue.Size())
	log.Println()
	log.Println()

	for i := 0; i < 30; i++ {
		deQueue, e := queue.DeQueue()
		if e != nil {
			log.Println(e)
			continue
		}
		log.Println(deQueue)
	}

	log.Println()
	log.Println()
	log.Println(queue.IsEmpty())
	log.Println(queue.IsFull())
	log.Println(queue.Size())
	log.Println()

	for i := 0; i < 100; i++ {
		deQueue, e := queue.DeQueue()
		if e != nil {
			log.Println(e)
			continue
		}
		log.Println(deQueue)
	}

	log.Println()
	log.Println()
	log.Println(queue.IsEmpty())
	log.Println(queue.IsFull())
	log.Println(queue.Size())
	log.Println()

	for i := 0; i < 30; i++ {
		e := queue.EnQueue("is good" + strconv.Itoa(i))
		if e != nil {
			log.Println(e)
			continue
		}
	}

	log.Println(queue.IsEmpty())
	log.Println(queue.IsFull())
	log.Println(queue.Size())

	queue.Clear()

	log.Println(queue.IsEmpty())
	log.Println(queue.IsFull())
	log.Println(queue.Size())
}
