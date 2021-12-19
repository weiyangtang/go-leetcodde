package interface_learn

import (
	"fmt"
	"testing"
)

type Printer interface {
	print()
}

type S struct{}

func (s S) print() {
	fmt.Println("interface")
}

func TestInterface(t *testing.T) {
	i := S{}
	i.print()
}

type Inter interface {
	Ping()
	Pong()
}

type Anter interface {
	Inter
	String()
}

type Str struct{}

func (Str) Ping() {
	fmt.Println("ping")
}

func (Str) Pong() {
	fmt.Println("pong")
}

func TestAssert(t *testing.T) {
	str := Str{}
	var i interface{} = str
	o := i.(Inter)
	o.Ping()
	o.Pong()
	fmt.Printf("%p,%p,%p\n", &str, &o, &i)

	if o, ok := i.(Inter); ok {
		fmt.Printf("%p\n", &o)
	}

	switch i.(type) {
	case Inter:
		fmt.Println("Inter type")
	case Anter:
		fmt.Println("Anter type")
	}

	input := 0

	switch {
	case input == 0:
		fmt.Println("0")
	case input-1 == -1:
		fmt.Println("1")
	default:
		fmt.Println("default")
	}
	u := "strconv.Atoi()"[0]
	r := rune(u)
	r2 := r - '0'
	if r2 <= 9 && r2 >= 0 {

	}
}
