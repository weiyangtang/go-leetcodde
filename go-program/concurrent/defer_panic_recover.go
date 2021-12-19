package concurrent

import "log"

func panicReachDefer() {
	log.Println("====start==== ")
	defer log.Println("defer func1")
	defer log.Println("defer func2")
	defer func() {
		v := recover()
		log.Println("after panic ", v)
	}()
	panic("myPanic")
	log.Println("throw panic")
}

type Runnable = interface {
	Run()
}
