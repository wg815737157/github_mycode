package main

import (
	"fmt"
)

func singleNumber(a []int) (int, int) {
	tmp := 0
	for i := 0; i < len(a); i++ {
		tmp ^= a[i]
	}

	bit := 1
	for i := 0; i < 32; i++ {
		if bit&tmp != 1 {
			bit <<= 1
		}
	}

	e1, e2 := 0, 0

	for i := 0; i < len(a); i++ {
		if bit&a[i] == bit {
			e1 ^= a[i]
		} else {
			e2 ^= a[i]
		}
	}
	return e1, e2
}

func main() {
	a := []int{1, 2, 3, 3, 5, 5}
	fmt.Println(singleNumber(a))
}
