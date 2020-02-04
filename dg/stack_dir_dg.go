package dg

import (
	"datastructure/stack"
	"io/ioutil"
	"log"
	filepath2 "path/filepath"
)

//  通过任务stack 实现目录遍历

func GetDirAllByStack(filepath string) ([]string,error) {
	stack := stack.New(200)
	filepath,err := filepath2.Abs(filepath)
	if err != nil {
		return nil,err
	}
	stack.Push(filepath)

	result := []string{}

	for !stack.IsEmpty() {
		pop, e := stack.Pop()
		if e != nil {
			log.Fatalln(e)
		}
		filepath = pop.(string)
		infos, e := ioutil.ReadDir(filepath)
		if e != nil {
			log.Println(e)
			continue
		}
		for _, v := range infos {
			if v.IsDir() {
				filepath = filepath2.Join(filepath,v.Name())
				result = append(result,filepath)
				stack.Push(filepath)
			}else {
				file := filepath2.Join(filepath,v.Name())
				result = append(result,file)
			}
		}
	}
	return result,nil
}
