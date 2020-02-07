package main

import (
	"fmt"
	"strings"
)

func bubbleSortStr(data []string) []string {
	for i := len(data) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if strings.Compare(data[j], data[j+1]) == -1 {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	return data
}

func main() {
	var testData = []string{"D", "o", "l", "l", "a", "r", "K", "i", "l", "l", "e", "r", "a", "s", "d", "j", "d", "g", "u", "r", "e", "h", "g", "u", "i", "r", "u", "i", "g", "u", "i", "e", "r", "f"}
	fmt.Println(bubbleSortStr(testData))
}
