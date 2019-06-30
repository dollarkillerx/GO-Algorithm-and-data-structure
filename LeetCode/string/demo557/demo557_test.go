/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-29
* Time: 上午10:05
* */
package demo557

import (
	"fmt"
	"strings"
	"testing"
)

func TestDemo557(t *testing.T) {
	data := "Let's take LeetCode contest"
	fmt.Println(reverseWords(data))
}

func reverseWords(s string) string {
	data := ""
	split := strings.Split(s, " ")
	end := len(split)-1
	for k,v := range split {
		i := desc(v)
		if k != end {
			data += i + " "
		}else{
			data += i
		}
	}
	return data
}

func desc(s string) string {
	data := ""
	for i:=len(s)-1;i>=0;i-- {
		data += string(s[i])
	}
	return data
}