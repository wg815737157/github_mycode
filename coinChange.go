package main

import "fmt"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] {
				dp[i] = twoMinNum(dp[i-coins[j]]+1, dp[i])
				fmt.Println(i, dp[i])
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
func myCoinChange(coins []int, amount int) int {
	res := make([]int, amount+1)
	for i, _ := range res {
		res[i] = -1
	}
	res[0] = 0
	for i := 1; i < amount+1; i++ {
		minValue := amount
		for _, v := range coins {
			if i-v < 0 || res[i-v] == -1 {
				continue
			}
			if res[i-v] < minValue {
				minValue = res[i-v]
			}
		}
		if minValue < amount {
			res[i] = minValue + 1
		}
	}
	return res[amount]
}
func main() {
	fmt.Println(coinChange([]int{122, 112, 383, 404, 25, 368}, 6977))
}
