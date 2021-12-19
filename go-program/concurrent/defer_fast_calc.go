package concurrent

import "log"

// defer 延迟函数会预先计算调用函数传入的实参，预先拷贝一份实参的变量

func deferComputeArgs() {
	for i := 0; i < 3; i++ {
		var i = i
		log.Println("loop:", i)
		defer func() {
			log.Println("a:", i)
		}()
	}
	for i := 0; i < 3; i++ {
		log.Println("b:", i)
	}
}
