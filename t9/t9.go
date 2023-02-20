package main

import (
	"fmt"
	"time"
)

func generator(channel_1, exitChan chan int, arr []int) {
	for _, val := range arr {
		channel_1 <- val
		time.Sleep(500 * time.Millisecond)
	}
}

func squarer(channel_1, channel_2, exitChan chan int) {
	for {
		select {
		case _, _ = <-exitChan:
			return
		default:
		}

		select {
		case x, ok := <-channel_1:
			if ok {
				channel_2 <- x * x
				continue
			} else {
				continue
			}
		default:
		}
	}
}

func printer(channel_2, exitChan chan int) {
	for {
		select {
		case _, _ = <-exitChan:
			return
		default:
		}

		select {
		case x, ok := <-channel_2:
			if ok {
				fmt.Printf("Signal %d was read!\n", x)
				continue
			} else {
				continue
			}
		default:
		}
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}

	channel_1 := make(chan int)
	channel_2 := make(chan int)
	exitChan := make(chan int)

	go generator(channel_1, exitChan, arr)
	go squarer(channel_1, channel_2, exitChan)
	go printer(channel_2, exitChan)

	time.Sleep(time.Second * 5)

	/* to kill go routines manually */

	// exitChan <- 1
	// exitChan <- 1

	// time.Sleep(time.Millisecond * 500)
}
