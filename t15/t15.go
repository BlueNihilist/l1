package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

var justString string = "abcdefghijklmnopqrtuvwxyz"

func createHugeString(size int) string {
	res := ""
	for i := 0; i < size; i++ {
		res += strconv.Itoa(i)
	}
	return res
}

func someFunc() {
	v := createHugeString(1 << 10)
	fmt.Printf("v = \"%s\"\n&v = %p, sizeof(v) = %d, len(v) = %d\n\n", v, &v, unsafe.Sizeof(v), len(v))

	justString = v[:100]
	fmt.Printf("justString = \"%s\"\n&justString = %p, sizeof(justString) = %d, len(jestString) = %d\n\n", justString, &justString, unsafe.Sizeof(justString), len(justString))

	// return justString
}

/* weeeiiiiird stuff is happening here */

func main() {
	someFunc()

	fmt.Printf("Shifting justString by 1 byte: ")

	// Сдвигаем указатель
	*(*uintptr)(unsafe.Pointer(&justString))++
	fmt.Printf("\"%s\"\n", justString)

	// Размер строкии не поменялся И мы одновременно видим кусок исходной строки v.
	// justString ссылается на исходую строку, которая не была удалена из памяти, так как на нее до сих пор кто-то указывает,
	// даже не смотря на то, что новая строка использует только кусок v.
	// Таким образом, мы теряем много памяти на ресурсы, которые не используются.
	// Для предотвращения этого можно, например, создавать новую строку через make, и с помощью цикла заносить туда нужные нам данные.

	fmt.Printf("len(justString) = %d\n", len(justString))
}
