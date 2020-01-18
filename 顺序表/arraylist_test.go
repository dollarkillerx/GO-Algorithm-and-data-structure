/**
*@program: GO-Algorithm-and-data-structure
*@description: https://github.com/dollarkillerx
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-01-17 19:55
 */
package array

import (
	"log"
	"testing"
)

func TestArray1(t *testing.T) {
	ar := make([]string, 10)
	log.Println(len(ar))
	log.Println(cap(ar))
	ar[0] = "acac"
	log.Println(ar)
}

func TestArrayList_Append(t *testing.T) {
	list := NewArrayList()
	list.Append("sadasd")
	list.Append("xsdadsad")
	list.Append("sadasd")
	list.Append("sadasd")
	list.Append("sadasd")
	list.Append("sadasd")
	list.Append("sadasd")
	list.Append("sadasd")
	list.Append("sadasd")
	list.Append("sadasd")
	list.Append("sadasd")
	list.Append("sadasd")
	list.Append("sadasd")
	list.Append("sadasd")
	list.Append("sadasd")
	log.Println(list.String())

	get, err := list.Get(1)
	if err != nil {
		log.Println(get)
	}
	_, err = list.Get(200)
	if err != nil {
		log.Println(err)
	}
	log.Println(list.size)

	err = list.Delete(1)
	err = list.Delete(1)
	err = list.Delete(1)
	err = list.Delete(1)
	err = list.Delete(1)
	err = list.Delete(1)
	err = list.Delete(1)
	err = list.Delete(1)
	err = list.Delete(1)
	err = list.Delete(1)
	err = list.Delete(1)
	err = list.Delete(1)
	err = list.Delete(1)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(list.String())

}
