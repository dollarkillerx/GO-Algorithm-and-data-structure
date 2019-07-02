/**
* Created by GoLand
* User: dollarkiller
* Date: 19-7-2
* Time: 上午10:16
* */
package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 返回测试数据
func getData() []int {
	var data []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 30; i++ {
		data = append(data, rand.Intn(100))
	}
	return data
}

// 冒泡排序
func TestBubbling(t *testing.T) {
	data := getData()

	for i := len(data) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			tmp := data[j]
			if tmp > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	fmt.Println(data)
}
