package main

import "fmt"

func findGCD(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func main() {
	fmt.Println(findGCD(3, 4))
}
