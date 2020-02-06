package main

import "fmt"

var testData = []int{2, 5, 8673, 534, 89, 1, 23, 0, 0, 2, 344, 654, 77, 2, -1}

func sort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	for i := 0; i < len(testData)-1; i++ {
		min := testData[i]
		for j := i + 1; j < len(testData); j++ {
			d1 := testData[j]
			if min > d1 {
				testData[j] = min
				min = d1
			}
		}
		testData[i] = min
	}
	return data
}

func main() {
	a := sort(testData)
	fmt.Println(a)
}
