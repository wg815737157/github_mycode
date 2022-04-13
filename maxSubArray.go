package main

import "fmt"

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if nums[i] > maxSum {
			maxSum = nums[i]
		}
	}
	return maxSum
}

func maxSubArray2(nums []int) int {
	maxSum, sum := nums[0], nums[0]
	n := len(nums)
	for i := 1; i < n; i++ {
		if sum < 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}

func main() {
	a := []int{1, 2, 4, -2, 3, 6, -9}
	fmt.Println(maxSubArray(a))
}
