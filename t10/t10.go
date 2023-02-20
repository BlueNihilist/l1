package main

import (
	"fmt"
)

func main() {
	arr := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -0.1, -9.4, 5.5, 9.1}
	m := make(map[int][]float32)
	for _, val := range arr {
		var key int
		if val < 0 {
			key = int((val-10)/10) * 10
		} else {
			key = int(val/10) * 10
		}
		m[key] = append(m[key], val)

		/* Если по примеру, то: */
		// key := int(val/10) * 10
		// m[key] = append(m[key], val)
	}
	fmt.Printf("%v\n", m)
}
