package main

import "fmt"

func findMin(a []int) int {
	n := len(a)
	for i := 0; i < n-1; i++ {
		if a[i] > a[i+1] {
			return a[i+1]
		}
	}
	return a[0]
}
func findMin2(a []int) int {
	n := len(a)
	l := 0
	r := n - 1
	for l < r {
		pivot := l + (r-l)/2
		if a[pivot] < a[r] {
			r = pivot
			continue
		}
		l = pivot + 1
	}
	return a[l]
}

func main() {
	arrayA := []int{1, 2, 3, 4}
	fmt.Println(findMin2(arrayA))
}
