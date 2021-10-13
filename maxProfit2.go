package main

import (
	"fmt"
	"math"
)

func maxProfile2(prices []int) int {
	n := len(prices)
	profile := make([][2]int, n)
	profile[0][0] = 0
	profile[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		profile[i][0] = twoMaxNum(profile[i-1][0], profile[i-1][1]+prices[i])
		profile[i][1] = twoMaxNum(profile[i-1][1], profile[i-1][0]-prices[i])
	}
	return profile[n-1][0]
}
func maxProfile3(prices []int) int {
	n := len(prices)
	profile := make([]int, 2)
	profile[0] = 0
	profile[1] = -prices[0]
	for i := 1; i < n; i++ {
		profile[0] = twoMaxNum(profile[0], profile[1]+prices[i])
		profile[1] = twoMaxNum(profile[1], profile[0]-prices[i])
	}
	return profile[0]
}

func maxProfile(prices []int) int {
	n := len(prices)
	if n < 2 {
		return -1
	}
	profile, minPrice := 0, math.MaxInt32

	for i := 0; i < n; i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
			continue
		}

		if prices[i]-minPrice > profile {
			profile = prices[i] - minPrice
		}
	}
	return profile
}
func main() {
	prices := []int{1, 3, 2, 4, 9, 3}
	fmt.Println(maxProfile(prices))
	fmt.Println(maxProfile2(prices))
	fmt.Println(maxProfile3(prices))
}
