package main

import "fmt"

func rotate2(a []int, k int) {
	n := len(a)
	m := k % n
	partA := (a)[:n-m]
	partB := (a)[n-m:]
	tmp := append(partB, partA...)
	copy(a, tmp)
}

func rotate(a []int, k int) {
	revertArray(a)
	k = k % len(a)
	revertArray(a[:k])
	revertArray(a[k:])
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(a, 3)
	fmt.Println(a)
}
