package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"time"
)
var res [][]int

func swap(intSlice []int, i, j int) {
	intSlice[i], intSlice[j] = intSlice[j], intSlice[i]
}

func iteration(intSlice []int, curPosition int) {
	l := len(intSlice)
	if curPosition == l {
		desSlice := make([]int, l)
		copy(desSlice, intSlice)
		res = append(res, desSlice)
		return
	}
	for i := curPosition; i < l; i++ {
		swap(intSlice, curPosition, i)
		iteration(intSlice, curPosition+1)
		swap(intSlice, curPosition, i)
	}
}

func permutation(intSlice []int) {
	iteration(intSlice, 0)
}

func permutation(num []int, n int, curIndex int, res *[][]int) {
	if curIndex == n-1 {
		fmt.Println(num)
		tmp := make([]int, len(num))
		copy(tmp, num)
		*res = append(*res, tmp)
		return
	}
	for i := curIndex; i < n; i++ {
		swap(num, curIndex, i)
		permutation(num, n, curIndex+1, res)
		swap(num, i, curIndex)
	}
}
