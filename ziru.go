package main

import (
	"fmt"
	"math"
)

var result [][]int

func solution(a []int, cur int) {
	math.MaxInt32
	if len(a) == cur {
		tmp := make([]int, len(a))
		copy(tmp, a)
		result = append(result, tmp)
		return
	}
	for i := cur; i < len(a); i++ {
		swap(a, cur, i)
		solution(a, cur+1)
		swap(a, i, cur)
	}
}

func main() {
	a := []int{1, 2, 3}
	solution(a, 0)
	fmt.Println(result)
}
