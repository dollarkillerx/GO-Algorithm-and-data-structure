package dg

import (
	"fmt"
	"log"
	"testing"
)

func TestGetDirAllByLinkStack(t *testing.T) {
	strings, e := GetDirAllByLinkStack("/home/dollarkiller/Github/Golang")
	if e != nil {
		log.Fatalln(e)
	}
	for _, v := range strings {
		fmt.Println(v)
	}
}
