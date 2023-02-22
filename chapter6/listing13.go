package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter1 int64
	awg      sync.WaitGroup
)

func list13() {
	awg.Add(2)

	go incCounter1(1)
	go incCounter1(2)

	awg.Wait()
	fmt.Println("Final Counter:", counter1)
}

func incCounter1(id int64) {
	defer awg.Done()

	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter1, 1)
		runtime.Gosched()
	}
}
