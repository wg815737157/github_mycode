package main

import "fmt"

func singleNumberInArray(input []int) []int {
	abXOR := 0
	for i := 0; i < len(input); i++ {
		abXOR ^= input[i]
	}

	div := 1
	for (div & abXOR) == 0 {
		div <<= 1
	}

	a, b := 0, 0

	for i := 0; i < len(input); i++ {
		if div&input[i] == div {
			a ^= input[i]
			continue
		}
		b ^= input[i]
	}
	return []int{a, b}
}
func main() {
	intput := []int{1, 2, 5, 2}
	fmt.Println(singleNumberInArray(intput))
}
