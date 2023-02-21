package main

import "fmt"

func main() {
	a := 5
	b := 10

	fmt.Printf("a = %d \tb = %d\n", a, b)

	a, b = b, a

	fmt.Printf("a = %d \tb = %d\n", a, b)

	a = a + b // A = a + b; B = b
	b = a - b // A = a + b; B = a
	a = a - b // A = a; B = b

	fmt.Printf("a = %d \tb = %d\n", a, b)
}
