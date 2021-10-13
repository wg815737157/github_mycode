package main

import (
	"fmt"
)

func minNum(a, b, c int) int {
	if a > b {
		if b > c {
			return c
		}
		return b
	}
	if a > c {
		return c
	}
	return a
}

func minerInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1, amount+1)
	for i := 1; i < amount+1; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i < amount+1; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] > i {
				continue
			}
			dp[i] = minerInt(dp[i], dp[i-coins[j]]+1)
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func main() {
	a := []int{2, 5, 10, 1}
	fmt.Println(coinChange(a, 27))

}
