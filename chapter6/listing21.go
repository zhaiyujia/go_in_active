package main

import (
	"fmt"
	"sync"
	"time"
)

var wg21 sync.WaitGroup

func list21() {
	baton := make(chan int)
	wg21.Add(1)

	go Runner(baton)

	baton <- 1

	wg21.Wait()
}

func Runner(baton chan int) {
	var newRunner int

	runner := <-baton

	fmt.Printf("Runner %d Running With Baton\n", runner)

	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", newRunner)
		go Runner(baton)
	}

	time.Sleep(100 * time.Millisecond)

	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		wg21.Done()
		return
	}

	baton <- newRunner
}
