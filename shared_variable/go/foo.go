// Use `go run foo.go` to run your program

package main

import (
	"fmt"
	"runtime"
)

var i = 0

func main() {
	// Setter at gorutinen kan bruke to operativsystemer samtidig.
	runtime.GOMAXPROCS(2)

	decrementChan := make(chan int)
	incrementChan := make(chan int)

	doneChan := make(chan bool, 2)

	fmt.Println("halloooo")

	go incrementing(incrementChan, doneChan)
	go decrementing(decrementChan, doneChan)

	fmt.Println("hei hei")

	choice(incrementChan, decrementChan, doneChan)

	//time.Sleep(500 * time.Millisecond)
	fmt.Println("The magic number is:", i)
}

func incrementing(incrementChan chan int, doneChan chan bool) {
	for j := 0; j < 100000000; j++ {
		incrementChan <- 1
	}
	close(incrementChan)
	doneChan <- true
	fmt.Println("Thread 1 ferdig")
}

func decrementing(decrementChan chan int, doneChan chan bool) {
	for j := 0; j < 100000000; j++ {
		decrementChan <- 1
	}
	close(decrementChan)
	doneChan <- true
	fmt.Println("Thread 2 ferdig")
}

func choice(incrementChan chan int, decrementChan chan int, doneChan chan bool) {
	doneCount := 0

	for {
		select {
		case val, ok := <-incrementChan:
			if ok {
				i += val
			}
		case val, ok := <-decrementChan:
			if ok {
				i -= val
			}
		case <-doneChan:
			doneCount++
			if doneCount == 2 {
				return
			}
		}
	}
}
