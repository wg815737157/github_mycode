package main

import (
	"fmt"
)



func quickSort2(input []int) {
	if len(input) <= 1 {
		return
	}
	pivot := input[len(input)-1]

	r := len(input) - 1
	i := 0

	for j := 0; j < r; j++ {
		if input[j] < pivot {
			i++
			swap(input, i, j)
		}
	}
	swap(input, i, r)
	quickSort2(input[:i])
	quickSort2(input[i+1:])
}

func quickSort(input []int) {
	if len(input) <= 1 {
		return
	}
	low := 0
	high := len(input) - 1
	pivot := input[0]
	for low < high {
		if input[high] >= pivot {
			high--
			continue
		}
		if input[low] <= pivot {
			low++
			continue
		}
		swap(input, low, high)
	}
	swap(input, low, 0)

	quickSort(input[:low])
	quickSort(input[low+1:])
	return
}

func main() {

	input := []int{5, 2, 3, 1}

	quickSort2(input)
	fmt.Println(input)
}
