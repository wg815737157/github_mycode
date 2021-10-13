package main

import "fmt"

func moveZeroes(nums []int) {
	n := len(nums)
	j := 0
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
}
func main() {
	a := []int{0, 1, 1, 2, 3, 45, 6, 9}
	moveZeroes(a)
	fmt.Println(a)
}
