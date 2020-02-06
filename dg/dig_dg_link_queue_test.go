package dg

import (
	"fmt"
	"testing"
)

func TestGetDirAllByLinkQueue(t *testing.T) {
	strings, e := GetDirAllByLinkQueue("/home/dollarkiller/Github/Golang")
	if e != nil {
		panic(e)
	}
	for _, v := range strings {
		fmt.Println(v)
	}
}
