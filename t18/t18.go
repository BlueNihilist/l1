package main

import (
	"fmt"
	"sync"
	"time"
)

type counter struct {
	val int
	mx  sync.Mutex
	ch  chan bool
}

func conc_1(cnt *counter) {
	for {
		cnt.ch <- true
		cnt.val += 1
		fmt.Printf("conc_1 set val = %d\n", cnt.val)
		<-cnt.ch
		time.Sleep(time.Millisecond * 100)
	}
}

func conc_2(cnt *counter) {
	for {
		time.Sleep(time.Millisecond * 110)
		cnt.mx.Lock()
		cnt.val += 1
		fmt.Printf("conc_2 set val = %d\n", cnt.val)
		cnt.mx.Unlock()
	}
}

func main() {
	var cnt counter
	cnt.ch = make(chan bool, 1)

	go conc_1(&cnt)
	go conc_2(&cnt)

	time.Sleep(time.Second)

	fmt.Printf("val = %d\n", cnt.val)
}
