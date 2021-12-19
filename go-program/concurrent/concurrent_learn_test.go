package concurrent

import (
	"log"
	"math/rand"
	"runtime"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

func SayGreetings(greeting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(greeting)
		d := time.Second * time.Duration(rand.Intn(5)) / 2
		time.Sleep(d)
	}
	wg.Done()
}

func TestCurrent01(t *testing.T) {
	rand.Seed(time.Now().Unix())
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
	log.Println("cpu:", runtime.NumCPU())
	wg.Add(2)
	go SayGreetings("hello tangweiyang", 5)
	go SayGreetings("hi liuni", 5)
	time.Sleep(2 * time.Second)
	wg.Wait()
}

func TestGoroutineDefer(t *testing.T) {
	complexDefer()
}

func TestDeferComputeArgs(t *testing.T) {
	deferComputeArgs()
}

func TestPanicReachDefer(t *testing.T) {
	panicReachDefer()
}
