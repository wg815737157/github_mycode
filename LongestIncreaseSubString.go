package main

import "fmt"

func LongestIncreaseSubString(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{nums}
	}
	dp := make([][]int, len(nums))
	dp[0] = append(dp[0], nums[0])
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			dpJMax := dp[j][len(dp[j])-1]
			dpJLen := len(dp[j])
			if nums[i] > dpJMax && len(dp[i]) < dpJLen {
				dp[i] = make([]int, len(dp[j]))
				copy(dp[i], dp[j])
			}
		}
		dp[i] = append(dp[i], nums[i])
	}
	return dp
}

func main() {
	nums := []int{2, 1, 8, 3, 7}
	res := LongestIncreaseSubString(nums)
	fmt.Println(res)
}
