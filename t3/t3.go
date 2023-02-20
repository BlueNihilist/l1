package main

import (
	"fmt"
	"sync"
)

func square(x int) int {
	return x * x
}

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	sum := 0
	wg := sync.WaitGroup{}
	for _, val := range arr {
		wg.Add(1)
		go func() {
			sum += square(val)
			wg.Done()
		}()
		wg.Wait()
	}
	fmt.Printf("Sum = %d\n", sum)
}
