package dg

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestGetDirAll(t *testing.T) {
	strings, e := GetDirAll("/home/dollarkiller/Github/Golang/GO-Algorithm-and-data-structure")
	if e != nil {
		log.Fatalln(e)
	}

	for _,v := range strings {
		log.Println(v)
	}
}

func TestRead(t *testing.T) {
	infos, e := ioutil.ReadDir("/home/dollarkiller/Github/Golang/GO-Algorithm-and-data-structure/dg")
	if e != nil {
		log.Fatalln(e)
	}
	log.Println(infos)
}