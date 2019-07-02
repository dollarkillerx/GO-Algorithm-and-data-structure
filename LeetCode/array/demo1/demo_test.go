/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-28
* Time: 下午12:34
* */
package demo1

import (
	"fmt"
	"testing"
)

func TestOne(t *testing.T) {
	nums := []int{1, 1, 2}
	removeDuplicates(nums)
	fmt.Println(nums)
}

func removeDuplicates(nums []int) int {
	ns := []int{}
	for _, v := range nums {
		if excts(ns, v) {
			ns = append(ns, v)

		}
	}
	nums = ns[:]
	fmt.Println(ns)
	return len(ns)
}

func excts(ns []int, val int) bool {
	for _, v := range ns {
		if v == val {
			return false
		}
	}
	return true
}
