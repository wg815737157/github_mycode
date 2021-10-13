package main

import "fmt"

func findRepeatNumber(a []int) int {
	mapA := make(map[int]bool)
	for i := 0; i < len(a); i++ {
		if _, ok := mapA[a[i]]; ok {
			return a[i]
		}
		mapA[a[i]] = true
	}
	return -1
}
func main() {
	setA := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(findRepeatNumber(setA))
}

