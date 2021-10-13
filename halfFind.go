package main

import "fmt"

func halfFind(a []int, target int) bool {
	low := 0
	high := len(a) - 1
	for low <= high {
		mid := low + (high-low)/2
		if a[mid] == target {
			return true
		}
		if a[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

func main() {
	a := []int{1, 2, 3, 5, 6, 7, 8, 9, 10}
	fmt.Println(halfFind(a, 9))
}
