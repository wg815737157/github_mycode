package main

import "fmt"

func hammingWeight(num uint32) int {
	res := 0
	detector := uint32(1)
	for i := 0; i < 32; i++ {
		if num&detector == detector {
			res++
		}
		detector <<= i
	}
	return res
}

func main() {
	a := -1
	fmt.Println(hammingWeight(uint32(a)))
}
