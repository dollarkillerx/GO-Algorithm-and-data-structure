/**
* Created by GoLand
* User: dollarkiller
* Date: 19-7-2
* Time: 上午11:09
* */
package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 返回测试数据
func getDatas(num int, max int) []int {
	var data []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < num; i++ {
		data = append(data, rand.Intn(max))
	}
	return data
}

func TestDemo164(T *testing.T) {
	data := getDatas(6, 99)
	fmt.Println(data)

	fmt.Println(maximumGap(data))

}

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	ints := sor(nums)

	max := ints[1] - ints[0]

	for i := 1; i < len(ints); i++ {
		ds := ints[i] - ints[i-1]
		if max < ds {
			max = ds
		}
	}

	return max
}

// 排序
func sor(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		min := nums[i]
		for j := len(nums) - 1; j > i; j-- {
			if min > nums[j] {
				min, nums[j] = nums[j], min
			}
		}
		nums[i] = min
	}
	return nums
}
