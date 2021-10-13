package main

import "fmt"

func plusOne(digits []int) []int {
	n := len(digits)
	carry := 1
	for i := n - 1; i >= 0; i-- {
		if digits[i]+carry == 10 {
			digits[i] = 0
		} else {
			digits[i]++
			return digits
		}
	}
	tmp := []int{1}
	digits = append(tmp, digits...)
	return digits
}

func main() {
	a := []int{2, 9, 1}
	fmt.Println(plusOne(a))
}
