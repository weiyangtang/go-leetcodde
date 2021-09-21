package reference

import (
	"fmt"
	"sort"
	"testing"
)

type BigObject struct {
	//data [100000000]byte
	val int
	//orderMap map[string]int
}

func input(in BigObject) BigObject {
	fmt.Printf("%v", in)
	in.val = 20
	return in
}

func TestReference(t *testing.T) {
	a := BigObject{val: 10}
	b := a
	fmt.Println(a == b)
	fmt.Println(input(a) == a)

}

func TestMapGet(t *testing.T) {
	hash := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println(hash["a"])
	list := []string{"a", "dd", "cc", "ee"}
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})
	fmt.Printf("%v", list)
}
