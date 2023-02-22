package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

type Action struct {
	Human
	Working bool
}

func (h *Human) printHuman() {
	fmt.Printf("Name: \"%s\" Age: %d\n", h.Name, h.Age)
}

func main() {
	var a Action

	a.Name = "Joe"
	a.Age = 81
	a.Working = true

	a.printHuman()
}
