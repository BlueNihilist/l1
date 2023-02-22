package main

import (
	"fmt"
	"time"
)

func sleep(dt time.Duration) {
	start := time.Now()
	for {
		if time.Now().Sub(start).Nanoseconds() > dt.Nanoseconds() {
			break
		}
	}
}

func main() {
	fmt.Println("Current date and time is: ", time.Now().String())
	fmt.Printf("Sleeping...\n")
	sleep(time.Second * 5)
	fmt.Println("Current date and time is: ", time.Now().String())
}
