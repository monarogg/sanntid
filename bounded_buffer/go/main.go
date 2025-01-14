package main

import (
	"fmt"
	"time"
)

// skal legge til data i bufferen
func producer(buffer chan int) {

	for i := 0; i < 10; i++ {
		buffer <- i
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("[producer]: pushing %d\n", i)
	}

}

// skal hente ut data fra bufferen
func consumer(buffer chan int) {

	time.Sleep(1 * time.Second)
	for {
		i := <-buffer
		fmt.Printf("[consumer]: %d\n", i)
		time.Sleep(50 * time.Millisecond)
	}

}

func main() {

	buffer := make(chan int, 5)

	go consumer(buffer)
	go producer(buffer)

	select {}
}
