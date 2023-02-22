package main

import "fmt"

func mirror(str string) string {
	arr := []rune(str)
	n := len(arr)
	for i := 0; i < (n / 2); i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
	return string(arr)
}

func main() {
	str := "главрыба"
	fmt.Printf("String: \"%s\" \nMirrored: \"%s\"\n\n", str, mirror((str)))
	str = "Любовь долготерпит, милосердствует, любовь не завидует, любовь не превозносится, не гордится, не бесчинствует, не ищет своего, не раздражается, не мыслит зла, не радуется неправде, а сорадуется истине; все покрывает, всему верит, всего надеется, все переносит. Любовь никогда не перестает, хотя и пророчества прекратятся, и языки умолкнут, и знание упразднится."
	fmt.Printf("String: \"%s\" \nMirrored: \"%s\"\n\n", str, mirror((str)))
	str = "Hello, 世界"
	fmt.Printf("String: \"%s\" \nMirrored: \"%s\"\n\n", str, mirror((str)))
}
