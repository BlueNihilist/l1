package main

import (
	"fmt"
	"time"
)

var noSeconds = 3

func ViaChannel(exitChan chan bool) {
	for {
		select {
		case x, ok := <-exitChan:
			if ok {
				fmt.Printf("Signal %t was read! Exiting ViaChanne()\n", x)
				return
			} else {
				fmt.Printf("Channel closed! Exiting ViaChannel()\n")
				return
			}
		default:
			fmt.Printf("ViaChannel() running\n")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func killViaChannel() {
	exitChan := make(chan bool)
	go ViaChannel(exitChan)
	time.Sleep(time.Duration(noSeconds) * time.Second)
	exitChan <- true
	close(exitChan)
}

func killViaGlobalVar() {
	exitFlag := false

	ViaGlobalVar := func() {
		for {
			if exitFlag {
				fmt.Printf("Flag is equal to %t! Exiting ViaGlobalVar()\n", exitFlag)
				return
			}
			fmt.Printf("ViaGlobalVar() running\n")
			time.Sleep(500 * time.Millisecond)
		}
	}

	go ViaGlobalVar()

	time.Sleep(time.Duration(noSeconds) * time.Second)
	exitFlag = true
}

func ViaPointer(exitFlag *bool) {
	for {
		if *exitFlag {
			fmt.Printf("Flag is equal to %t! Exiting ViaPointer()\n", *exitFlag)
			return
		}
		fmt.Printf("ViaPointer() running\n")
		time.Sleep(500 * time.Millisecond)
	}
}

func killViaPointer() {
	exitFlag := false

	go ViaPointer(&exitFlag)

	time.Sleep(time.Duration(noSeconds) * time.Second)
	exitFlag = true
}

func killViaExitFromMain() {
	go func() {
		for {
			fmt.Printf("killViaExitFromMain() running\n")
			time.Sleep(500 * time.Millisecond)
		}
	}()
}

func main() {
	killViaChannel()

	killViaGlobalVar()

	killViaPointer()

	killViaExitFromMain()

	time.Sleep(time.Duration(noSeconds) * time.Second)
}
