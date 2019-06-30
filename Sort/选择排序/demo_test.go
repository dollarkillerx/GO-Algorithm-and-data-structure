/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-21
* Time: 下午10:30
* */
package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestOne(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	data := []int{}
	for i:=0;i<100;i++{
		intn := rand.Intn(999)
		data = append(data, intn)
	}
	fmt.Println(data)

	len := len(data)
	//end := len - 1
	s := 0
	for s<len{
		min := data[s]
		st := s
		for ;st<len;st++{
			if data[st] < min {
				min,data[st] = data[st],min
			}
		}
		data[s] = min
		s += 1 // 每移动一遍 就向前推动一格
	}
	fmt.Println(data)
}