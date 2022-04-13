package main

import "fmt"

func halfMultiply(a int, n int) int {
	if n == 1 {
		return a
	}
	res := 0
	if n%2 == 0 {
		res = halfMultiply(a*a, n/2)
	} else {
		res = halfMultiply(a*a, n/2) * a
	}
	return res
}

func main() {
	fmt.Println(halfMultiply(2, 8))
}

