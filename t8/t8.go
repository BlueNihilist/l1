package main

import (
	"errors"
	"fmt"
)

func setIthBitTo0(x int64, i int) (int64, error) {
	if i > 63 || i < 0 {
		return x, errors.New("Incorrect i value!")
	}
	rhs := int64(-1)
	rhs = rhs ^ (1 << i)
	return x & rhs, nil
}

func setIthBitTo1(x int64, i int) (int64, error) {
	if i > 63 || i < 0 {
		return x, errors.New("Incorrect i value!")
	}
	rhs := int64(1 << i)
	return x | rhs, nil
}

func main() {
	x := int64(1)
	fmt.Printf("x = \t\t\t %d\n", x)

	x, _ = setIthBitTo1(x, 4)
	fmt.Printf("Set 4th bit to 1:\t %d\n", x)

	x, _ = setIthBitTo1(x, 63)
	fmt.Printf("Set 63rd bit to 1:\t %d\n", x)

	x, _ = setIthBitTo0(x, 63)
	fmt.Printf("Set 63rd bit to 0:\t %d\n", x)

	x, _ = setIthBitTo0(x, 4)
	fmt.Printf("Set 4th bit to 0:\t %d\n", x)
}
