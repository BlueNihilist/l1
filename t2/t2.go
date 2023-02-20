package main

import (
	"fmt"
	"sync"
)

func square(x int) {
	fmt.Printf("%d * %d = %d\n", x, x, x*x)
}

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	for _, val := range arr {
		wg.Add(1)
		go func() {
			square(val)
			wg.Done()
		}()
		wg.Wait()
	}
}
