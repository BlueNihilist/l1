package main

import (
	"fmt"
	"math/rand"
	"time"
)

func readChannel(channel chan int) {
	for {
		fmt.Printf("Read: %d\n", <-channel)
		time.Sleep(500 * time.Millisecond)
	}
}

func writeChannel(channel chan int) {
	for {
		channel <- rand.Intn(100)
	}
}

func main() {
	noSeconds := 5

	channel := make(chan int, 10)
	defer close(channel)

	go writeChannel(channel)
	go readChannel(channel)

	time.Sleep(time.Duration(noSeconds) * time.Second)
}
