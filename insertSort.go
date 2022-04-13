package main

import "fmt"

//循环 实现方法 插排
func insertSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for j := 1; j < len(nums); j++ {
		key := nums[j]
		i := j - 1
		for ; i >= 0 && nums[i] > key; i-- {
			nums[i+1], nums[i] = nums[i], nums[i+1]
		}
		nums[i+1] = key
	}
}

func main() {
	nums := []int{4, 2, 7, 3, 2}
	insertSort(nums)
	fmt.Println(nums)

}
