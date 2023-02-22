package main

import (
	"fmt"
	"time"
)

func square(x int) int {
	return x * x
}

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	sum := 0
	ch := make(chan bool, 1)

	for _, val := range arr {
		go func(x int) {
			ch <- true
			sum += square(x)
			<-ch
		}(val)
	}

	fmt.Printf("Sum = %d\n", sum)
	time.Sleep(time.Second)
	fmt.Printf("Sum = %d\n", sum)
}
