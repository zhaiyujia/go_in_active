package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	lwg     sync.WaitGroup
)

func list09() {
	lwg.Add(2)

	go incCounter(1)
	go incCounter(2)
	lwg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incCounter(id int) {
	defer lwg.Done()

	for count := 0; count < 2; count++ {
		value := counter

		// 当前goro 从县城推出，并刚回到队列
		runtime.Gosched()

		value++

		counter = value
	}
}
