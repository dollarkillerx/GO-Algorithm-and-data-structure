package main

import (
	"fmt"
)

func bubbleSortNum(data []int) []int {
	a := 0
	for i := len(data) - 1; i > 0; i-- {
		isok := false
		a++
		for j := 0; j < i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				isok = true
			}
		}
		if !isok {
			break
		}
	}
	fmt.Println(a)
	return data
}

func main() {
	var testData = []int{23, 1, 34, 45, 6, 3, 765, 7, 786, 8, 67, 843, 42, 4, 234, 0, -2}
	//var testData = []int{0,1,2,3,4,5,6,7,8,9,10}
	num := bubbleSortNum(testData)
	fmt.Println(num)
}
