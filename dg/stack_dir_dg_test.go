package dg

import (
	"fmt"
	"log"
	"testing"
)

func TestGetDirAllByStack(t *testing.T) {
	strings, e := GetDirAllByStack("/home/dollarkiller/Github/Golang/GO-Algorithm-and-data-structure")
	if e != nil {
		log.Fatalln(e)
	}
	for _, v := range strings {
		fmt.Println(v)
	}
}
