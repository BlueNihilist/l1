package main

import (
	"fmt"
	"strings"
)

func reverseSlice(arr []string) []string {
	n := len(arr)
	for i := 0; i < (n / 2); i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
	return arr
}

func reverseOrderOfWords(str string) string {
	words := strings.Split(str, " ")
	words = reverseSlice(words)
	return strings.Join(words, " ")
}

func main() {
	str := "Hello Words!"
	fmt.Printf("String: \"%s\" \nReversed: \"%s\"\n\n", str, reverseOrderOfWords((str)))
	str = "你好 谷歌 翻譯"
	fmt.Printf("String: \"%s\" \nReversed: \"%s\"\n\n", str, reverseOrderOfWords((str)))
	str = "И познаете истину, и истина сделает вас свободными"
	fmt.Printf("String: \"%s\" \nReversed: \"%s\"\n\n", str, reverseOrderOfWords((str)))
}
