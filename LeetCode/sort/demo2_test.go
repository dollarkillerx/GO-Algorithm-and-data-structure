/**
* Created by GoLand
* User: dollarkiller
* Date: 19-7-2
* Time: 上午10:43
* */
package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 返回测试数据
func getData2() []int {
	var data []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 30; i++ {
		data = append(data, rand.Intn(100))
	}
	return data
}

// 选择排序
func TestDemo2(t *testing.T) {
	data2 := getData2()
	fmt.Println(data2)

	for i:=0;i<len(data2);i++{
		min := data2[i]
		for i2:=len(data2)-1;i2>i;i2--{
			if min > data2[i2] {
				min,data2[i2] = data2[i2],min
			}
		}
		data2[i] = min
	}

	fmt.Println(data2)
}
