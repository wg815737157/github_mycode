package main

import "fmt"

func firstMissingPositive2(nums []int) int {
	l := len(nums)
	mapA := make(map[int]bool)
	for i := 0; i < l; i++ {
		mapA[nums[i]] = true
	}
	i := 1
	for ; i <= l+1; i++ {
		if _, ok := mapA[i]; ok {
			continue
		}
		return i
	}
	return i
}



func firstMissingPositive(nums []int) int {

	l := len(nums)
	for i := 0; i < l; i++ {
		if nums[i] <= 0 {
			nums[i] = l + 1
		}
	}

	for i := 0; i < l; i++ {
		tmp := abs(nums[i])
		if tmp <= l {
			nums[tmp-1] = -abs(nums[tmp-1])
		}
	}
	for i := 0; i < l; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return l + 1
}
func main() {
	fmt.Println(firstMissingPositive([]int{1, 1}))
}
