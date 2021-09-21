package stringPrac

import (
	"fmt"
	"math"
	"testing"
	"unicode/utf8"
	"unsafe"
)

func TestString(t *testing.T) {
	s1 := "我很开心吗？"
	for i := 0; i < len(s1); i++ {
		fmt.Println(string(s1[i]))
	}
}

func TestUtf8String(t *testing.T) {
	s := "Hello, 中国"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
	fmt.Println(utf8.RuneSelf)
	fmt.Println(math.Pow(2, 16) - 1)
	/*
		utf8 字符编码
		0xxxxxxx  文字符号（rune） 1-127 ACSII
		110xxxxx 10xxxxxx 文字符号在 128~2047
		1110xxxx 10xxxxxx 10xxxxxx   2^11 ~ 2^16 -1
		11110xxx 10xxxxxx 10xxxxxx 10xxxxxx 2^16 ~ 2^21-1
	*/
	fmt.Println(0xF0)
	fmt.Println(15 * 16)
	fmt.Println(0b11110111 % 16)
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\t%d\n", i, r, size)
		i += size
	}

	for i, r := range s {
		fmt.Printf("%d\t%c\n", i, r)
	}
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d\t%c\n", i, s[i])
	}

	var strBytes []byte
	strBytes = append(strBytes, "hello"...)
	fmt.Printf("%s\n", strBytes)

}

func TestBytes2String(t *testing.T) {
	var bytes []byte = make([]byte, 10, 10)
	bytes[0] = 'a'
	bytes[1] = 'c'
	bytes[2] = 'd'
	var s string = string(bytes)
	fmt.Println(s)
	println(s)
	var s1 *string = (*string)(unsafe.Pointer(&bytes))
	s2 := s1
	fmt.Println(s1)
	fmt.Println(*s2)
}
