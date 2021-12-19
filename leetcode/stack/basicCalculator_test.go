package stack

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTokens(t *testing.T) {
	expression := "-15/12*(-14-8)+17"
	//expression := "-1*(-14-8)"
	tokens := getTokens(expression)
	fmt.Println(tokens)
}

func TestIsDigit(t *testing.T) {
	flag := isDigit('1')
	assert.Equal(t, true, flag)
	flag = isDigit('*')
	assert.Equal(t, false, flag)
}

func TestConvertRPN(t *testing.T) {
	rpn := convertRPN([]string{"-15", "/", "12", "*", "(", "-14", "-", "-8", ")", "+", "17"})
	fmt.Println(rpn)
}

func TestBasicCalculator(t *testing.T) {
	//expression := "-15/12*(-14-8)+17"
	//expression := "(1+(4+5+2)-3)+(6+8)"
	expression := "- (3 + (4 + 5))"
	res := calculate(expression)
	fmt.Println(res)
	//val := -15/12*(-14-8)+17
	val := -(3 + (4 + 5))
	fmt.Println(val)
}
