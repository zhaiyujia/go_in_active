package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	numberGoroutines = 4
	taskLoad         = 10
)

var wg24 sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func list24() {
	tasks := make(chan string, taskLoad)

	wg24.Add(numberGoroutines)

	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker24(tasks, gr)
	}

	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	close(tasks)

	wg24.Wait()
}

func worker24(tasks chan string, worker int) {
	defer wg24.Done()

	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("Worker: %d : Shutting Down\n", worker)
			return
		}

		fmt.Printf("Worker: %d : Started %s\n", worker, task)

		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}
}
