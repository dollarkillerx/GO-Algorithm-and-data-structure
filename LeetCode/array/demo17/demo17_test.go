/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-30
* Time: 上午9:16
* */
package demo17

import (
	"fmt"
	"testing"
)

func TestDemo17(t *testing.T) {
	data := []int{1,2,3}

	fmt.Println(data[2:])
	//data := []interface{} {"abc","abd","ghi"}
	//js(data)
}

func js(data []interface{}) {
	switch data[0].(type) {
	case string:
		rest := []string{}
		for _,v := range data[0].(string) {
			for _,v2 := range data[1].(string) {
				dat := string(v) + string(v2)
				rest = append(rest,dat)
				dats := data[1:]

			}
		}
	case []string:
		rest := []string{}
		for _,v := range data[0].([]string) {
			for _,v2 := range data[1].(string) {
				dat := string(v) + string(v2)
				rest = append(rest,dat)
				data[0] = rest
			}
		}
	}
	if len(data) > 1 {
		js(data)
	}else{
		fmt.Println(data)
	}
}
