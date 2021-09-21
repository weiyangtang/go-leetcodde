package defer_learn

import (
	"fmt"
	"testing"
	"time"
)

func trace() func() {
	fmt.Println("trace enter method")
	return trace2()
}

func trace2() func() {
	fmt.Println("trace2 enter method")
	return func() {
		fmt.Println("trace2 quit method")
	}
}

func deferLearn() {
	defer trace()()
	fmt.Println("hello world")
}

type DeferClass struct {
}

func (t *DeferClass) deferTest1() func() {
	fmt.Println("class method")
	return func() {
		fmt.Println("return s")
	}
}

func TestDefer(t *testing.T) {
	//deferTest2()
	fmt.Println("method start")
	class := DeferClass{}
	defer class.deferTest1()()
	//fn := class.deferTest1()
	//defer fn()
	fmt.Println("test defer")
}

func deferTest2() {
	startedAt := time.Now()
	//defer fmt.Println(time.Since(startedAt))
	defer func() func() {
		fmt.Println(time.Since(startedAt))
		return func() {
			fmt.Println("all finished")
		}
	}()
	fmt.Println("start sleep")
	time.Sleep(time.Second)
}

type number int

func (n number) print()   { fmt.Println(n) }
func (n *number) pprint() { fmt.Println(*n) }

func TestDefer3(t *testing.T) {
	var n number

	defer n.print()
	defer n.pprint()
	defer func() { n.print() }()
	defer func() { n.pprint() }()

	n = 3
}

func TestDeferFunc(t *testing.T) {
	deferLearn()
}

func TestPreCompute(t *testing.T) {
	startTime := time.Now()
	fmt.Println(startTime)
	defer fmt.Printf("startTime:%v, now:%v, cost time:%v", startTime, time.Now(), time.Since(startTime))
	time.Sleep(10 * time.Second)
}
