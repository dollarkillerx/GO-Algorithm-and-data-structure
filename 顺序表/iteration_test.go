/**
*@program: GO-Algorithm-and-data-structure
*@description: https://github.com/dollarkillerx
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-01-18 10:19
 */
package array

import (
	"log"
	"testing"
)

func TestArrayList_Iteration(t *testing.T) {
	list := NewArrayList()
	list.Append("1232")
	list.Append("1232")
	list.Append("1232")
	list.Append("1232")
	list.Append("1232")
	list.Append("1232")

	for iteration := list.Iteration(); iteration.IsNext(); {
		next, err := iteration.Next()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(next)
	}
}
