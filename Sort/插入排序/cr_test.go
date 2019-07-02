/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-22
* Time: 上午10:55
* */
package cr

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestCRPX(t *testing.T) {
	data := make([]int, 10)
	rand.Seed(time.Now().UnixNano())
	for i := range data {
		data[i] = rand.Intn(999)
	}
	fmt.Println(data)

	// 插入排序
	for i := range data {
		for j := i; j > 0; j-- {
			if data[j] < data[j-1] {
				data[j], data[j-1] = data[j-1], data[j]
			}
		}
	}
	fmt.Println(data)
}
