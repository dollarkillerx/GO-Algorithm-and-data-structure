/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-21
* Time: 上午9:49
* */
package demo

import (
	"testing"
)
// 已知a+b+c=1000,且a^2+b^2=c^2,却1000内所有a,b,c
func TestDemo1(t *testing.T) {
	// 外层循环遍历A
	for a:=0;a<3001;a++ {

		// 循化得到B   此时时间复杂度为n^2
		for b:=0;b<3001;b++ {

			// 时间复杂度n^3
			for c:=0;c<3001;c++ {
				if a + b + c == 1000 && a*a+b*b == c*c {
					t.Logf("找到一个:a:%d ,b:%d ,c:%d\n",a,b,c)
				}
			}

		}
	}
}

func TestDemo2(t *testing.T) {
	for a:=0;a<3001;a++{
		for b:=0;b<3001;b++{
			c := 1000-b-a
			if a*a+b*b == c*c{
				t.Logf("找到一个:a:%d ,b:%d ,c:%d\n",a,b,c)
			}
		}
	}
}


