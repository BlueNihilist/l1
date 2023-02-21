package main

import "fmt"

func quick_sort(arr *[]int, indFirst, indLast int) {
	if indFirst >= indLast {
		return
	}

	indLeft := indFirst
	indRight := indLast
	mid := (*arr)[(indLeft+indRight)/2]

	loopBody := func() {
		for (*arr)[indLeft] < mid {
			indLeft += 1
		}
		for (*arr)[indRight] > mid {
			indRight -= 1
		}
		if indLeft <= indRight {
			(*arr)[indLeft], (*arr)[indRight] = (*arr)[indRight], (*arr)[indLeft]
			indLeft += 1
			indRight -= 1
		}
	}

	for {
		loopBody()
		if indLeft > indRight {
			break
		}
	}

	quick_sort(arr, indFirst, indRight)
	quick_sort(arr, indLeft, indLast)

}

func main() {
	arr := []int{4, 1, 2, 2, 9, 1, -4, 0, 0, -9, 19, 3}
	fmt.Printf("unsorted arr = %v\n", arr)
	quick_sort(&arr, 0, len(arr)-1)
	fmt.Printf("sorted arr =   %v\n", arr)
}
