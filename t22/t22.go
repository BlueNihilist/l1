package main

import (
	"fmt"
	"math/big"
)

func main() {
	x_1 := big.NewInt(1 << 30)
	fmt.Printf("x_1 = %v\n", x_1)

	x_2 := big.NewInt(1 << 62)
	fmt.Printf("x_2 = %v\n", x_2)

	x_3 := big.NewInt(0)
	x_3.Mul(x_2, x_1)
	fmt.Printf("x_3 = x_2 * x_1 = %v\n", x_3)

	x_4 := big.NewInt(0)
	x_4.Div(x_2, x_1)
	fmt.Printf("x_4 = x_2 / x_1 = %v\n", x_4)

	x_5 := big.NewInt(0)
	x_5.Sub(x_2, x_1)
	fmt.Printf("x_5 = x_2 - x_1 = %v\n", x_5)

	x_6 := big.NewInt(0)
	x_6.Add(x_1, x_2)
	fmt.Printf("x_6 = x_2 + x_1 = %v\n", x_6)
}
