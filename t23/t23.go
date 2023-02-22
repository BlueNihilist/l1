package main

import (
	"errors"
	"fmt"
)

func delIthElem(slice []int, i int) ([]int, error) {
	if i < 0 || i >= len(slice) {
		return slice, errors.New(fmt.Sprintf("Index i = %d outside of range [%d, %d]", i, 0, len(slice)-1))
	}
	return append(slice[:i], slice[i+1:]...), nil
}

func main() {
	slice := []int{0, 1, 2, 3, 4, 5}
	fmt.Printf("Slice = %v\n", slice)

	slice, _ = delIthElem(slice, 2)
	fmt.Printf("Slice = %v\n", slice)

	slice, _ = delIthElem(slice, 0)
	fmt.Printf("Slice = %v\n", slice)

	slice, _ = delIthElem(slice, 2)
	fmt.Printf("Slice = %v\n", slice)

	slice, _ = delIthElem(slice, 2)
	fmt.Printf("Slice = %v\n", slice)

	tmpSlice, err := delIthElem(slice, 2)
	if err != nil {
		fmt.Printf("delIthElem: %v\n", err)
	} else {
		fmt.Printf("Slice = %v\n", tmpSlice)
	}
}
