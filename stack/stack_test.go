package stack

import (
	"fmt"
	"log"
	"testing"
)

func TestStack_New(t *testing.T) {
	stack := New(10)

	fmt.Println(stack.IsEmpty())
	fmt.Println(stack.IsFull())

	for i := 0; i < 10; i++ {
		err := stack.Push("this is good")
		if err != nil {
			log.Println(err)
		}
	}

	fmt.Println(stack.Size())

	for i:=0;i<10;i++ {
		pop, e := stack.Pop()
		if e != nil {
			log.Println(e)
		}
		s1,bool := pop.(string)
		if bool {
			fmt.Println(s1)
		}else {
			fmt.Println("is not")
		}
	}

	//ar := stack.GetAR()
	//fmt.Println(ar)
	//pop, e := stack.Pop()
	//if e != nil {
	//	fmt.Println(e)
	//}
	//fmt.Println(pop)

	log.Println(stack.Size())
	//pop, e = stack.Pop()
	//if e != nil {
	//	log.Fatalln(e)
	//}
}
