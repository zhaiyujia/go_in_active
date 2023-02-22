package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg20 sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func list20() {

	court := make(chan int)
	go player("Nadal", court)
	go player("Djokovic", court)
	court <- 1
	wg20.Add(2)

	wg20.Wait()
}

func player(name string, court chan int) {
	defer wg20.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s Won\n", name)
			return
		}

		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			close(court)
			return
		}
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++
		court <- ball
	}

}
