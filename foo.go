// Use `go run foo.go` to run your program

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var i = 0
var mu sync.Mutex

func main() {
	// Setter at gorutinen kan bruke to operativsystemer samtidig.
	runtime.GOMAXPROCS(2)

	go incrementing()
	go decrementing()

	time.Sleep(500 * time.Millisecond)
	fmt.Println("The magic number is:", i)
}

func incrementing() {
	for j := 0; j < 1000000; j++ {
		mu.Lock()
		i++
		mu.Unlock()
	}
}

func decrementing() {
	for j := 0; j < 1000000; j++ {
		mu.Lock()
		i--
		mu.Unlock()
	}
}

// func choice(incrementChan chan int, decrementChan chan int, doneChan chan bool) {
// 	doneCount := 0

// 	for {
// 		select {
// 		case val, ok := <-incrementChan:
// 			if ok {
// 				i += val
// 			}
// 		case val, ok := <-decrementChan:
// 			if ok {
// 				i -= val
// 			}
// 		case <-doneChan:
// 			doneCount++
// 			if doneCount == 2 {
// 				return
// 			}
// 		}
// 	}
// }
