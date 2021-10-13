package main

import "fmt"

func swapHeap(array []int, i int, size int) {
	lager := i
	l, r := i*2+1, i*2+2
	if l < size && array[i] < array[l] {
		lager = l
	}
	if r < size && array[i] < array[r] {
		lager = r
	}
	if lager != i {
		array[i], array[lager] = array[lager], array[i]
		swapHeap(array, lager, size)
	}
}

func myBuildMaxHeap(array []int, size int) []int {
	for i := size/2 - 1; i >= 0; i-- {
		swapHeap(array, i, size)
	}
	return array
}
func kthLargest(array []int, k int) int {
	if k >= len(array) {
		return -1
	}
	len := len(array)
	for i := 0; i < k; i++ {
		myBuildMaxHeap(array, len-i)
		array[0], array[len-1-i] = array[len-1-i], array[0]
	}
	return array[len-k]
}

func main() {
	array := []int{3, 6, 12, 4, 7, 45}
	res := kthLargest(array, 2)
	fmt.Println(res)
}
