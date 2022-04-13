package main

import "fmt"

func maxProfit(prices []int) int {
	maxProfit := 0
	if len(prices) < 2 {
		return maxProfit
	}
	min := prices[0]
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			if prices[i]-min > maxProfit {
				maxProfit = prices[i] - min
			}
		}
	}
	return maxProfit
}
func main() {
	fmt.Println(maxProfit([]int{1, 2, 7, 4, 5, 7, 9, 3}))
}
