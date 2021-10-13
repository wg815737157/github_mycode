package main

import (
	"fmt"
)

func sortedSquares(nums []int) []int {

	n := len(nums)
	res := make([]int, n)
	i := 0
	j := n - 1
	pos := n - 1
	for i <= j {
		tmp1 := nums[i] * nums[i]
		tmp2 := nums[j] * nums[j]
		if tmp1 > tmp2 {
			res[pos] = tmp1
			i++
		} else {
			res[pos] = tmp2
			j--
		}
		pos--
	}
	return res
}

func main() {
	a := []int{-9, -5, -4, 0, 1, 2, 3, 4}
	fmt.Println(sortedSquares(a))
}
