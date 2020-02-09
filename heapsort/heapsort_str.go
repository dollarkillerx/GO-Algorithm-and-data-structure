package main

import "fmt"

func buildSortStr(data []string) {
	laskNode := len(data) - 1
	head := (laskNode - 1) / 2
	for i := head; i >= 0; i-- {
		heapSortStr(data, len(data), i)
	}
}

func heapSortStr(data []string, len int, i int) {
	if len <= i {
		return
	}
	k1 := 2*i + 1
	k2 := 2*i + 2
	max := i
	if k1 < len && data[k1] > data[max] {
		k1 = max
	}

	if k2 < len && data[k2] > data[max] {
		k2 = max
	}
	if max != i {
		data[max], data[i] = data[i], data[max]
		heapSortStr(data, len, max)
	}
}

func sortStr(data []string) {
	buildSortStr(data)
	fmt.Println(data)
	for i:= len(data) -1;i>0;i--{
		data[0],data[i] = data[i],data[0]
		heapSortStr(data,len(data),i)
	}
}

func main() {
	data := []string{"a","s","d","f","g","h","j","k","l","z","x","c","v","b","b","n","M","q","w","e","r","t","Y","U","I","O","p"}
	sortStr(data)
	fmt.Println(data)

}
