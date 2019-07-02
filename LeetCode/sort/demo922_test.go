/**
* Created by GoLand
* User: dollarkiller
* Date: 19-7-2
* Time: 上午11:59
* */
package sort

import (
	"fmt"
	"testing"
)

//func TestDemo(t *testing.T) {
//	a := []int{1,2,3,4,5,6}
//	a = a[1:]
//	fmt.Println(a)
//}

//func TestDemo922(t *testing.T) {
//	data := []int{4,2,5,7,8,9}
//
//	var a []int // 奇数
//	var b []int // 偶数
//
//	for _,v := range data {
//		if v % 2 == 0 {
//			b = append(b,v)
//		}else if v != 0{
//			a = append(a,v)
//		}
//	}
//
//	for k,_ := range data {
//		switch  {
//		case k == 0:
//			data[k] = b[0]
//			b = append(b[1:])
//		case k % 2 == 0:
//			data[k] = b[0]
//			b = append(b[1:])
//		case k % 2 != 0:
//			data[k] = a[0]
//			a = append(a[1:])
//		}
//	}
//
//	fmt.Println(data)
//}

// 方法二
func TestDemo922Pro(t *testing.T) {
	data := []int{4, 2, 5, 7, 8, 9}

	for i := 0; i < len(data); i++ {
		if i == 0 || i%2 == 0 {
			if data[i]%2 != 0 {

				for j := len(data) - 1; j > i; j-- {
					if data[j]%2 == 0 {
						data[i], data[j] = data[j], data[i]
					}
				}

			}
		} else {
			if data[i]%2 == 0 {

				for j := len(data) - 1; j > i; j-- {
					if data[j]%2 != 0 {
						data[i], data[j] = data[j], data[i]
					}
				}

			}
		}

	}

	fmt.Println(data)
}
