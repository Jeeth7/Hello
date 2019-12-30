package main

import (
	"fmt"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var logCh2 = make(chan logEntry, 50)
var doneCh = make(chan struct{}) // This is a signal only channel of type empty struct

//Ways of closing channel

func main() {
	go logger()
	go logger2()
	defer func() {
		close(logCh) // need to close the channel
	}()
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}
	logCh2 <- logEntry{time.Now(), logInfo, "App is starting"}
	logCh <- logEntry{time.Now(), logInfo, "App shutting down"}
	logCh2 <- logEntry{time.Now(), logInfo, "App shutting down"}

	doneCh <- struct{}{} // pass on the empty struch to signal the closure of the channel
	time.Sleep(100 * time.Millisecond)
}

func logger() { //you
	for entry := range logCh {
		fmt.Printf("(1)%v - [%v]%v\n", entry.time.Format("2006-01-02:[15:04:05.000]"), entry.severity, entry.message)
	}
}

func logger2() { //you
	for {
		select {
		case entry := <-logCh2:
			fmt.Printf("(2)%v - [%v]%v\n", entry.time.Format("2006-01-02:[15:04:05.000]"), entry.severity, entry.message)
		case <-doneCh:
			fmt.Println("done with chan 2")
			break
			//default: Having a default will make this method a non blocking
		}

	}
}
