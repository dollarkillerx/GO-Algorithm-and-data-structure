/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-19
* Time: 下午12:58
* */
package main

import (
	"errors"
	"log"
)

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i, e := que(data, 9)
	if e != nil {
		log.Println(e.Error())
	} else {
		log.Println(i)
	}
}

func que(data []int, m int) (int, error) {
	s := 0
	e := len(data)
	for s <= e {
		gg := (s + e) / 2
		i := data[gg]
		if i == m {
			return i, nil
		} else if i > m {
			e = i - 1
		} else {
			s = i + 1
		}
	}
	return 0, errors.New("no")
}
