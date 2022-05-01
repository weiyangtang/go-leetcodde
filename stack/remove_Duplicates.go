package stack

import (
	"fmt"
	"reflect"
)

func removeDuplicates(s string) string {
	stack := make([]byte, 0)
	for i, c := range s {
		fmt.Println(reflect.TypeOf(c))
		fmt.Println(reflect.TypeOf(s[i]))
		if len(stack) == 0 || stack[len(stack)-1] != s[i] {
			stack = append(stack, s[i])
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return string(stack)
}
