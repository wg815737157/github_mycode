package main

import (
	"fmt"
	"math"
)

func aContainB(a []int, target []int) bool {
	mapA := make(map[int]int)
	for i := 0; i < len(a); i++ {
		if _, ok := mapA[a[i]]; !ok {
			mapA[a[i]] = 0
		}
		mapA[a[i]]++
	}
	for i := 0; i < len(target); i++ {
		if _, ok := mapA[target[i]]; !ok {
			return false
		}
		mapA[target[i]]--
		if mapA[target[i]] < 0 {
			return false
		}
	}
	return true
}
func solution(a []int, target []int) int {
	n := len(a)
	i, j := 0, len(target)-1
	maxLen := math.MaxInt32
	for j <= n {
		//包含
		if aContainB(a[i:j], target) {
			if len(a[i:j]) < maxLen {
				maxLen = len(a[i:j])
			}
			i++
			continue
		}
		//不包含
		if !aContainB(a[i:j], target) {
			j++
		}
	}
	if maxLen == math.MaxInt32 {
		return -1
	}
	return maxLen
}
func main() {
	a := []int{1, 2, 3, 4, 1, 1, 4, 5, 3, 2}
	fmt.Println(solution(a, []int{1, 1, 2, 3, 4, 5}))
}
