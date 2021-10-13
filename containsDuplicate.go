package main

import "fmt"

func containsDuplicate(a []int) bool {
	mapA := make(map[int]bool)
	for i := 0; i < len(a); i++ {
		if _, ok := mapA[a[i]]; ok {
			return true
		}
		mapA[a[i]] = true
	}
	return false
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(containsDuplicate(a))
}
