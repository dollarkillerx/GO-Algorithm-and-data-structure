package link_stack

import (
	"fmt"
	"strconv"
	"testing"
)

func TestLinkStack_Clear(t *testing.T) {
	stack := New()

	fmt.Println(stack.IsEmpty())
	fmt.Println(stack.Size())

	for i := 0; i < 100; i++ {
		stack.Push("is go " + strconv.Itoa(i))
	}

	fmt.Println(stack.IsEmpty())
	fmt.Println(stack.Size())

	for !stack.IsEmpty() {
		pop, e := stack.Pop()
		if e != nil {
			fmt.Println(e)
			continue
		}
		fmt.Println(pop)
	}

	fmt.Println(stack.IsEmpty())
	fmt.Println(stack.Size())

	for i := 0; i < 100; i++ {
		stack.Push("is go " + strconv.Itoa(i))
	}

	stack.Clear()
	fmt.Println(stack.IsEmpty())
	fmt.Println(stack.Size())

}
