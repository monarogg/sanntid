package main

import (
	"fmt"
	"sync"
)

// 3: SHARING A VARIABLE

var i int
var mu sync.Mutex
var wg sync.WaitGroup

func thread_1() {
	defer wg.Done()
	for j := 0; j < 100000000; j++ {
		mu.Lock()
		i++
		mu.Unlock()
	}
	fmt.Println("Thread 1 ferdig")
}

func thread_2() {
	defer wg.Done()
	for j := 0; j < 100000000; j++ {
		mu.Lock()
		i--
		mu.Unlock()
	}
	fmt.Println("Thread 2 ferdig")
}

func main() {
	wg.Add(2)
	go thread_1()
	go thread_2()

	wg.Wait()

	fmt.Println("Resultat:", i)
}

// 4: SHARING A VARIABLE BUT PROPERLY
