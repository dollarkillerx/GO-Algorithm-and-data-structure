/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-30
* Time: 下午5:07
* */
package demo914

import (
	"fmt"
	"testing"
)

//func TestSer(t *testing.T)  {
//	s := 8
//	i := s % 2
//	fmt.Println(i)
//}

func TestDemo914(t *testing.T) {
	data1 := []int{1, 2, 3, 4, 4, 3, 2, 1}
	data2 := []int{1, 1, 1, 2, 2, 2, 3, 3}

	fmt.Println(hasGroupsSizeX(data1))
	fmt.Println(hasGroupsSizeX(data2))
}

func hasGroupsSizeX(deck []int) bool {
	for _, v := range deck {
		b := esc(v, deck)
		if b == false {
			return false
		}
	}
	return true
}

func esc(i int, ds []int) bool {
	var s int
	for _, v := range ds {
		if v == i {
			s++
		}
	}
	i2 := s % 2
	if i2 == 0 {
		return true
	}
	return false
}
