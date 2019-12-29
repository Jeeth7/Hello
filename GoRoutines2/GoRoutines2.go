package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int = 0
var wg = sync.WaitGroup{}
var m = sync.RWMutex{}

func main() {
	runtime.GOMAXPROCS(100)                             // TO set the number of threads
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1)) // runtime.GOMAXPROCS(-1) returs how many threads have been set
	for i := 0; i < 10; i++ {
		// just for sample on RLOCK and Lock this is a bad  practices
		wg.Add(2)
		m.RLock() // Read lock. allows multiple ready but no write
		go sayHello()
		m.Lock() // Write lock.
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Printf("\nHello #%v", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
