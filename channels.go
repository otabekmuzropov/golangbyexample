package main

import (
	"fmt"
	"time"
)

// Overview
// Channel is a data type in Go which provides synchrounization
// and communication between goroutines. They can be thought of as pipes
// which is used by goroutines to communicate.
// This communication between goroutines doesn’t require any explicit locks.
// Locks are internally managed by channel themselves. Channel along with goroutine
// makes the go programming language concurrent. So we can say that golang has two  concurrency
// primitives:

// Goroutine – lightweight independent execution to achieve concurrency/parallelism.
// Channels – provides synchronization and communication between goroutines.
// There are two major operations which can be done on a channel
// 	-Send
// 	-Receive

func main() {
	ch := make(chan int)

	fmt.Println("Sending value to channel")
	go send(ch)

	fmt.Println("Receiving from channel")
	go receive(ch)

	time.Sleep(time.Second)
}

func send(ch chan int) {
	// for i := 0; i < 7; i++ {
	// 	ch <- i
	// }
	ch <- 1
}

func receive(ch chan int) {
	val := <-ch
	fmt.Printf("Value Received=%d in receive function\n", val)
}
