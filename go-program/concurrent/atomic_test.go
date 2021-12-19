package concurrent

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestAtomicLoad(t *testing.T) {
	var commonInt32 int32 = 0
	go func() {
		//loadValue := atomic.LoadInt32(&commonInt32)
		commonInt32 = 1
		//time.Sleep(10 * time.Second)
		//fmt.Println(loadValue)
		fmt.Println(commonInt32)
		math.Ceil(float64(commonInt32))
	}()
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println(time.Now())
		//startTime := time.Now()
		//loadValue := atomic.LoadInt32(&commonInt32)
		loadValue := commonInt32
		fmt.Printf("time:%v,loadValue:%v\n", time.Now(), loadValue)
	}()
	time.Sleep(20 * time.Second)
}
