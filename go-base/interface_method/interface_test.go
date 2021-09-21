package interface_method

import (
	"fmt"
	"testing"
)

type Duck interface {
	Quack()
}

type Cat struct{}

func (c *Cat) Quack() {
	fmt.Println("meow")
}

func TestMethod(t *testing.T) {
	//cat := Cat{}
	//cat.Quack()
	var c = &Cat{}
	(c).Quack()
}
