/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-21
* Time: 上午10:48
* */
package demo

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

//func TestDem1(t *testing.T) {
//	n:= 1000
//	// 第一部部分时间复杂度为 O(n)
//	for i:=0;i<n;i++ {
//		fmt.Println(i)
//	}
//	// 第二部分时间复杂度为O(n^2)
//	for i:=0;i<n;i++ {
//		for b:=0;b<n;i++ {
//			fmt.Println(i)
//		}
//	}
//}

//func TestDem2(t *testing.T) {
//	n:=1000
//	flag := 1
//
//	// 第一部分时间复杂度O(n)
//	if flag >= 0 {
//		for a:=0;a<n;a++ {
//			fmt.Println(a)
//		}
//	// 第二部分时间复杂度O(n^2)
//	}else{
//		for a:=0;a<n;a++ {
//			for b:=0;b<n;b++ {
//				fmt.Println(b)
//			}
//		}
//	}
//}

func TestDasha(t *testing.T) {
	//// 模拟小红选的数
	//data := rand.Intn(101)
	//
	//// 大傻一个一个的猜
	//for sa:=0;sa<101;sa++ {
	//	if data == sa {
	//		fmt.Printf("大傻猜中了,猜了%d次\n",(sa+1))
	//	}
	//}
}

func TestErF(t *testing.T) {
	go func() {
		// 模拟小红选的数
		hong := rand.Intn(101)
		fmt.Println("hong:", hong)
		datas := make([]int, 0)
		for i := 0; i <= 100; i++ {
			datas = append(datas, i)
		}
		len := len(datas)
		fmt.Println(len)
		start := 0
		end := len - 1
		for start < end {
			// 中间数
			ms := (start + end) / 2
			fmt.Println("")
			if datas[ms] == hong {
				fmt.Printf("找到了hong:%d len:%d  data:%d\n", hong, ms, datas[ms])
				break
			} else if datas[ms] > hong {
				end = datas[ms]
				continue
			} else if datas[ms] < hong {
				start = datas[ms]
				continue
			}
		}
	}()
	time.Sleep(time.Second)
}
