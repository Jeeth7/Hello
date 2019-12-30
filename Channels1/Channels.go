package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan string)
	numChan := make(chan int)
	//byteChan := make(chan byte)

	fmt.Println("Hi Channels sample")
	wg.Add(2)
	go func() {
		fmt.Println("start")
		defer fmt.Println("stop")
		v := <-ch
		fmt.Println(v)
		v = <-ch
		fmt.Println(v)
		wg.Done()
	}()

	go func() {
		for nmch := range numChan {
			fmt.Println(nmch)
		}
		fmt.Println("I am done")
		wg.Done()
	}()

	go func(ch <-chan int) {
		for {
			if i, ok := <-ch; ok {
				fmt.Printf("\n - manual %v", i)
			} else {
				break
			}
		}
	}(numChan)
	time.Sleep(500 * time.Millisecond)
	ch <- "Hi this is a channel"
	ch <- "Sending in the same channel"
	//	wg.Done()

	numChan <- 3
	numChan <- 5
	numChan <- 1
	numChan <- 2
	numChan <- 7
	numChan <- 10
	close(numChan) // this will triger the go routines on range numchan to stop iterate otherwise deadlock

	wg.Wait()
}
