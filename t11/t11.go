package main

import (
	"fmt"
)

// func union(arr_1, arr_2 []int) []int {
// 	res := []int{}

// 	m := make(map[int]bool)

// 	for _, val := range arr_1 {
// 		m[val] = true
// 	}
// 	for _, val := range arr_2 {
// 		m[val] = true
// 	}

// 	for key := range m {
// 		res = append(res, key)
// 	}

// 	return res
// }

func intersection(arr_1, arr_2 []int) []int {
	res := []int{}

	m := make(map[int]int)

	for _, val := range arr_1 {
		m[val] += 1
	}
	for _, val := range arr_2 {
		m[val] += 1
	}

	for key, val := range m {
		if val == 2 {
			res = append(res, key)
		}
	}

	return res
}

func main() {
	set_1 := []int{1, 2, 3, 4, 5, 6, 7}
	set_2 := []int{2, 4, 6, 8, 10, 12, 14}

	fmt.Printf("%v\n", intersection(set_1, set_2))
}
