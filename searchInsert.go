package main

import "fmt"

func searchInsert(nums []int, target int) int {
	n := len(nums)
	high := n - 1
	low := 0

	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

func main() {
	a := []int{1, 2, 3, 5, 6, 8, 9}
	fmt.Println(searchInsert(a, 0))
}
