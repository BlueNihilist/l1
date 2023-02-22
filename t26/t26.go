package main

import (
	"fmt"
	"strings"
)

func lettersUnique(str string) bool {
	lowerStr := strings.ToLower(str)
	arr := []rune(lowerStr)
	m := make(map[rune]bool)

	for _, letter := range arr {
		if m[letter] {
			return false
		} else {
			m[letter] = true
		}
	}

	return true
}

func main() {
	strs := []string{"abcd", "abCdefAaf", "aabcd", "你好谷歌翻譯", "譯譯", "Русские буквы", "Ещё буквы"}
	for _, s := range strs {
		fmt.Printf("\"%s\" -- %t\n", s, lettersUnique(s))
	}
}
