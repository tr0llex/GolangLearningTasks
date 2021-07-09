package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Counter struct {
	Number int32
}

func (counter *Counter) AddOne() {
	atomic.AddInt32(&counter.Number, 1)
}

func main() {
	var counter Counter
	for i := 0; i < 1000; i++ {
		go counter.AddOne()
	}
	time.Sleep(time.Millisecond * 100)
	fmt.Println(counter)
}
