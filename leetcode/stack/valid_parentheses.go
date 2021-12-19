package stack

import "fmt"

func isValid(s string) bool {
	sym := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := make([]rune, 0)
	for _, c := range s {
		fmt.Println(c)
		if val, ok := sym[c]; !ok {
			stack = append(stack, c)
		} else {
			if val == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, val)
			}
		}
	}
	return len(stack) == 0
}
