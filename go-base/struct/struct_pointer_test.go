package _struct

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStructPointer(t *testing.T) {
	type Employee struct {
		Name   string
		Age    int
		Salary int
	}
	employee1 := &Employee{Name: "tangweiyang", Age: 10, Salary: 3000}
	employee2 := Employee{Name: "tangweiyang", Age: 10, Salary: 3000}
	fmt.Printf("%s\n%s\n", reflect.TypeOf(employee1), reflect.TypeOf(employee2))
	fmt.Println(employee1)
	fmt.Println(employee2)
	fmt.Println(employee1.Salary)
	fmt.Println(employee2.Name)
	fmt.Println(employee1 == nil)
	fmt.Println((*employee1).Name)
}
