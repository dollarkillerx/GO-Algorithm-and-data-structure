package main

import (
	"fmt"
)

// 生成堆
func buildHeap(data []int) {
	lastNode := len(data) - 1
	head := (lastNode - 1) / 2
	for i := head; i >= 0; i-- {
		heapSort(data, len(data), i)
	}
}

// 堆sort
func heapSort(data []int, len int, i int) {
	if len <= i {
		return
	}
	k1 := 2*i + 1
	k2 := 2*i + 2
	max := i
	if k1 < len && data[k1] > data[max] {
		max = k1
	}
	if k2 < len && data[k2] > data[max] {
		max = k2
	}
	if max != i {
		data[i], data[max] = data[max], data[i]
		heapSort(data, len, max)
	}
}

func heapSortGo(data []int) []int {
	buildHeap(data)
	it := len(data)
	for i := it - 1; i > 0; i-- {
		data[0], data[i] = data[i], data[0]
		heapSort(data, i, 0)
	}
	return data
}

func main() {
	//test1()
	//test2()
	test3()
}

func test3() {
	data := []int{0, 23, 4, 567, 568, 69, 2143, 25, 34, 656, 878, 9, 3, 4123, 211}
	data = heapSortGo(data)
	fmt.Println(data)
}

func test2() {
	data := []int{0, 23, 4, 567, 568, 69, 2143, 25, 34, 656, 878, 9, 3, 4123, 211}
	buildHeap(data)
	fmt.Println(data)
}

func test1() {
	//data := []int{4, 10, 3, 5, 1, 2}
	data := []int{4, 10, 3, 5, 1, 2}
	heapSort(data, len(data), 0)
	fmt.Println(data)
}
