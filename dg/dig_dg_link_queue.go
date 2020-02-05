package dg

import (
	"datastructure/link_queue"
	"io/ioutil"
	"log"
	"path/filepath"
)

func GetDirAllByLinkQueue(path string) ([]string, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	result := make([]string, 0)
	queue := link_queue.New()
	queue.EnQueue(path)
	for !queue.IsEmpty() {
		deQueue, err := queue.DeQueue()
		if err != nil {
			log.Println(err)
			continue
		}

		path, ok := deQueue.(string)
		if !ok {
			continue
		}

		infos, err := ioutil.ReadDir(path)
		if err != nil {
			log.Println(err)
			continue
		}
		for _, v := range infos {
			path := filepath.Join(path, v.Name())
			result = append(result, path)
			if v.IsDir() {
				queue.EnQueue(path)
			}
		}

	}
	return result, nil

}
