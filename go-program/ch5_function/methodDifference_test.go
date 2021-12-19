package ch5_function

import (
	"fmt"
	"testing"
)

type PointerInstance struct {
}

func (p *PointerInstance) pointer() {
	fmt.Printf("pointer:%p\n", p)
}

func (p PointerInstance) uPointer() {
	fmt.Printf("uPointer :%p\n", p)
}

func TestMethod(t *testing.T) {
	p := PointerInstance{}
	fmt.Printf("p :%p\n", &p)
	(*PointerInstance).pointer(&p)
	//PointerInstance.pointer(p)
	p.pointer()
	(&p).pointer()
}

func chPointer(a *int) {
	fmt.Printf("value:%p\n", a)
}

func TestPointerFunc(t *testing.T) {
	a := 10
	fmt.Printf("value:%p\n", &a)
	chPointer(&a)
}
