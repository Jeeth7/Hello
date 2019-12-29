package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	defer fmt.Println("\nFinish______________________")
	var ch string = "Hello"
	wg.Add(3)
	go func(msg *string) {
		fmt.Printf("\n1) %v", *msg)
		wg.Done()
	}(&ch)
	go func(msg string) {
		fmt.Printf("\n2) %v", msg)
		wg.Done()
	}(ch)
	go func() {
		fmt.Printf("\n3) %v", ch)
		wg.Done()
	}() // Race condition. not a good proactice
	ch = "Good morning"
	wg.Wait()
	// time.Sleep(time.Millisecond * 100) best practice is to use the wait group

	//Detect the race condtion at compile time
	// ***  go run -race filename.go ******
}
