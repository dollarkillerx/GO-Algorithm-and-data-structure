/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-19
* Time: 下午12:52
* */
package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestMans(t *testing.T) {
	data := make([]int, 100)
	rand.Seed(time.Now().UnixNano())
	for i := range data {
		data[i] = rand.Intn(999)
	}
	fmt.Println(data)
	Asort(data)

}

func Asort(data []int) {
	for i := range data {
		for j := 1; j < len(data)-i; j++ {
			if data[j] < data[j-1] {
				data[j], data[j-1] = data[j-1], data[j]
			}
		}
	}
	fmt.Println(data)
}
