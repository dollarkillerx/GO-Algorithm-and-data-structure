/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-30
* Time: 下午5:44
* */
package demo605

import (
	"fmt"
	"testing"
)

func TestDemo605(t *testing.T) {
	data := []int{1}
	flowers := canPlaceFlowers(data, 0)
	fmt.Println(flowers)
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	var i int
	if len(flowerbed) < 2 {
		if flowerbed[0] == 0 {
			i++
		}
		goto panduan
	}
	for k := 0; k < len(flowerbed); k++ {
		switch {
		case k == 0:
			if flowerbed[0] == 0 && flowerbed[1] == 0 {
				i++
				k++
			}
		case k != len(flowerbed)-1:
			if flowerbed[k] == 0 && flowerbed[k-1] == 0 && flowerbed[k+1] == 0 {
				i++
				k++
			}
		case k == len(flowerbed)-1:
			if flowerbed[k] == 0 && flowerbed[k-1] == 0 {
				i++
			}
		}
	}
panduan:
	if i >= n {
		return true
	}
	return false
}
