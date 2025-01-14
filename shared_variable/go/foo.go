// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
	"time"
)

var i = 0

func incrementing() {
	for j := 0; j < 100000000; j++ {
		i++
	}
	//TODO: increment i 1000000 times
}

func decrementing() {
	for j := 0; j < 100000000; j++ {
		i--
	}
	//TODO: decrement i 1000000 times
}

func main() {
	// Setter at gorutinen kan bruke to operativsystemer samtidig.
	runtime.GOMAXPROCS(2)

	go incrementing()
	go decrementing()

	// We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
	// We will do it properly with channels soon. For now: Sleep.
	time.Sleep(500 * time.Millisecond)
	Println("The magic number is:", i)
}
