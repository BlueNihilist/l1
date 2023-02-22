package main

import (
	"fmt"
	"time"
)

var noElems int = 100

func toChan(ch chan int) {
	for i := 0; i < noElems; i++ {
		ch <- i
		time.Sleep(time.Millisecond)
	}
}

func toMap(ch chan int) {
	m := make(map[int]int)
	for i := 0; i < 2*noElems; i++ {
		val := <-ch
		m[val%2] += val
	}
	fmt.Printf("Sum of evens: %d\n", m[0])
	fmt.Printf("Sum of odds: %d\n", m[1])
}

func main() {
	ch := make(chan int)

	go toChan(ch)
	go toChan(ch)
	go toMap(ch)

	time.Sleep(time.Second)
}
