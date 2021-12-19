package ch5_function

import (
	"fmt"
	"testing"
)

// 闭包：持有外部变量的函数

func f(base int) (func(int) int, func(int) int) {
	fmt.Println(&base, base)
	return func(i int) int {
			fmt.Println(&base, base)
			base += i
			return base
		}, func(i int) int {
			fmt.Println(&base, base)
			base -= i
			return base
		}
}

func TestClosure(t *testing.T) {
	addFunc1, subFunc1 := f(0)
	fmt.Println(addFunc1(1), subFunc1(2))

	addFunc2, subFunc2 := f(0)
	fmt.Println(addFunc2(1), subFunc2(2))
}

func TestBuildInFunc(t *testing.T) {
	slice := make([]int, 1, 10)
	slice = append(slice, 1)
}
