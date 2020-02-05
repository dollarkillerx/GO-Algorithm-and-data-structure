package dg

import (
	"datastructure/link_queue"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func GetDirAllByLinkStack(path string) ([]string, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	result := make([]string, 0)
	stack := link_queue.New()
	stack.EnQueue(path)

	for !stack.IsEmpty() {
		data1, err := stack.DeQueue()
		if err != nil {
			fmt.Println(err)
			continue
		}
		rs, ok := data1.(string)
		if !ok {
			continue
		}

		infos, err := ioutil.ReadDir(rs)
		if err != nil {
			fmt.Println(err)
			continue
		}
		for _, v := range infos {
			path := filepath.Join(rs, v.Name())
			result = append(result, path)
			if v.IsDir() {
				stack.EnQueue(path)
			}
		}
	}
	return result, nil
}
