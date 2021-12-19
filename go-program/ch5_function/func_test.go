package ch5_function

import (
	"fmt"
	"testing"
)

func TestIterate(t *testing.T) {
	//var slice []func()
	data := make([]int, 0)
	sli := []int{1, 2, 3, 4, 5}
	for _, v := range sli {
		fmt.Println(&v)
		data = append(data, v)
		//slice = append(slice, func() {
		//	fmt.Println(v * v) // 直接打印结果
		//})
	}

	//for _, val := range slice {
	//	val()
	//}

	for _, val := range data {
		fmt.Println(val)
	}
}
