package main

import "fmt"

func sortNum(data []int) []int {
	for i := 1; i < len(data); i++ {
		for j := i; j > 0; j-- {
			if data[j] < data[j-1] {
				data[j], data[j-1] = data[j-1], data[j]
			}
		}
	}
	return data
}

func main() {
	var testData = []int{23, 1, 34, 45, 6, 3, 765, 7, 786, 8, 67, 843, 42, 4, 234, 0, -2}
	num := sortNum(testData)
	fmt.Println(num)
}
