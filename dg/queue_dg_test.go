package dg

import (
	"fmt"
	"log"
	"testing"
)

func TestGetDirAllByQueue(t *testing.T) {
	path := "/home/dollarkiller/Github/Golang/GO-Algorithm-and-data-structure"
	strings, e := GetDirAllByQueue(path)
	if e != nil {
		log.Fatalln(e)
	}

	for _, v := range strings {
		fmt.Println(v)
	}
}
