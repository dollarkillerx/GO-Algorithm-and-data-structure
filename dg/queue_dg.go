package dg

import (
	"datastructure/queue"
	"io/ioutil"
	"log"
	"path/filepath"
)

func GetDirAllByQueue(path string) ([]string, error) {
	result := []string{}
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	queue := queue.New()
	queue.EnQueue(path)

	for !queue.IsEmpty() {
		s1 := queue.DeQueue()
		path, ok := s1.(string)
		if !ok {
			log.Fatalln("is err")
		}
		infos, err := ioutil.ReadDir(path)
		if err != nil {
			log.Println(err)
			continue
		}
		for _, v := range infos {
			if v.IsDir() {
				path := filepath.Join(path, v.Name())
				queue.EnQueue(path)
				result = append(result, path)
			} else {
				path := filepath.Join(path, v.Name())
				result = append(result, path)
			}
		}
	}
	return result, nil
}
