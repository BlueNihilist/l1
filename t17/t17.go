package main

import (
	"errors"
	"fmt"
)

func binary_search(arr []int, val, left, right int) (int, error) {
	if right < left {
		return 0, errors.New(fmt.Sprintf("Element %d is not present!", val))
	}

	mid := left + (right-left)/2

	if arr[mid] == val {
		return mid, nil
	}

	if arr[mid] > val {
		return binary_search(arr, val, left, mid-1)
	}

	return binary_search(arr, val, mid+1, right)
}

func main() {

	arr := []int{-3, 0, 2, 4, 6, 8, 11}

	val := 11
	ind, err := binary_search(arr, val, 0, len(arr)-1)
	if err == nil {
		fmt.Printf("Element %d is at index %d\n", val, ind)
	} else {
		fmt.Printf("%v\n", err)
	}

	val = -3
	ind, err = binary_search(arr, val, 0, len(arr)-1)
	if err == nil {
		fmt.Printf("Element %d is at index %d\n", val, ind)
	} else {
		fmt.Printf("%v\n", err)
	}

	val = 6
	ind, err = binary_search(arr, val, 0, len(arr)-1)
	if err == nil {
		fmt.Printf("Element %d is at index %d\n", val, ind)
	} else {
		fmt.Printf("%v\n", err)
	}

	val = 3
	ind, err = binary_search(arr, val, 0, len(arr)-1)
	if err == nil {
		fmt.Printf("Element %d is at index %d\n", val, ind)
	} else {
		fmt.Printf("%v\n", err)
	}

}
