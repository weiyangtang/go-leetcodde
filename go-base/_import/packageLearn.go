package _import

import "fmt"

var a = b + c
var b = globalVariable()
var c = 300

func globalVariable() int {
	return 3
}

func testVar() {
	fmt.Printf("%d %d %d\n", a, b, c)
}
