package dg

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func GetDirAll(filePath string) ([]string, error) {
	s, e := filepath.Abs(filePath)
	if e != nil {
		log.Println(1)
		return nil, e
	}
	fmt.Println("sc: ", s)

	infos, e := ioutil.ReadDir(filePath)

	if e != nil {
		log.Println(2)
		return nil, e
	}

	result := []string{}

	for _, v := range infos {
		bool := v.IsDir()
		if bool {
			filePath = filepath.Join(s, v.Name())
			result = append(result, filePath)
			strings, e := GetDirAll(filePath)
			if e != nil {
				log.Println(e)
				continue
			}
			result = append(result, strings...)
		}
	}

	return result, nil
}
