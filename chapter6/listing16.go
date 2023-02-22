package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter16 int
	wg16      sync.WaitGroup

	mutex sync.Mutex
)

func list16() {
	wg16.Add(2)

	go incCounter16(1)
	go incCounter16(2)

	wg16.Wait()
	fmt.Printf("Final Counter: %d", counter16)
}

func incCounter16(id int) {
	defer wg16.Done()

	for count := 0; count < 2; count++ {
		mutex.Lock()
		{
			value := counter16
			runtime.Gosched()
			value++
			counter16 = value
		}
		mutex.Unlock()
	}
}
