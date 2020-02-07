package main

import (
	"fmt"
	"strings"
)

//func getTesTData() []string  {
//	s1 := "DollarKillerasdjdgurehguiruiguierf"
//	fmt.Print("{")
//	for _,v := range s1 {
//		fmt.Printf(",\"%s\"",string(v))
//	}
//	fmt.Print("}")
//	return nil
//}

func sortString(testData []string) []string {
	for i := 0; i < len(testData)-1; i++ {
		min := testData[i]
		for j := i + 1; j < len(testData); j++ {
			a := testData[j]
			if strings.Compare(min, a) == 1 {
				//testData[j] = min
				//min = a
				testData[j], testData[i] = testData[i], testData[j]
			}
		}
		//testData[i] = min
	}
	return testData
}

func main() {
	//testDataStr()
	var testData = []string{"D", "o", "l", "l", "a", "r", "K", "i", "l", "l", "e", "r", "a", "s", "d", "j", "d", "g", "u", "r", "e", "h", "g", "u", "i", "r", "u", "i", "g", "u", "i", "e", "r", "f"}
	fmt.Println(sortString(testData))
}

func testDataStr() {
	fmt.Println(strings.Compare("a", "b")) // -1
	fmt.Println(strings.Compare("a", "a")) // 0
	fmt.Println(strings.Compare("A", "a")) // -1
	fmt.Println(strings.Compare("b", "a")) // 1
}
