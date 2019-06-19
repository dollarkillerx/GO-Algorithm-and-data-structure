/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-19
* Time: 下午12:52
* */
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var data [10]int
	for i:=0;i<10;i++{
		data[i] = rand.Intn(99)
	}
	fmt.Println(data)


}

func Asort(data []int) {

}