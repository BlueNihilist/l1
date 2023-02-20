package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, channel chan int) {
	for data := range channel {
		fmt.Printf("Worker #%d -- Message: %d\n", id, data)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	channel := make(chan int)

	N := 0
	fmt.Printf("Input Number of Workers: ")
	fmt.Scanf("%d", &N)

	for i := 0; i < N; i++ {
		go worker(i, channel)
	}

	for {
		data := rand.Intn(100)
		channel <- data
	}
}
