package main

import (
	"fmt"
	"reflect"
)

func printType(x interface{}) {
	fmt.Printf("Type: %v\n", reflect.TypeOf(x))
}

func main() {
	s := "string"
	i := 10
	f := 1.2
	b := true
	c := make(chan bool)

	fmt.Printf("var s = %s -- ", s)
	printType(s)

	fmt.Printf("var i = %d -- ", i)
	printType(i)

	fmt.Printf("var f = %f -- ", f)
	printType(f)

	fmt.Printf("var b = %t -- ", b)
	printType(b)

	fmt.Printf("var c = %v -- ", c)
	printType(c)
}
