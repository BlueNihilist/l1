package main

import (
	"fmt"
)

func toSet(strs []string) []string {
	res := []string{}

	m := make(map[string]bool)

	for _, val := range strs {
		m[val] = true
	}

	for key := range m {
		res = append(res, key)
	}

	return res
}

func main() {
	strs := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Printf("Set of strs: %v\n", toSet(strs))
}
