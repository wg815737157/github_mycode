package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quickSort(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	l, r := 0, 0
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Int31n(int32(n))
	nums[randomNum], nums[n-1] = nums[n-1], nums[randomNum]
	for r < n-1 {
		if nums[r] < nums[n-1] {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
		r++
	}
	nums[l], nums[n-1] = nums[n-1], nums[l]
	quickSort(nums[:l])
	quickSort(nums[l+1:])
}

func main() {
	input := []int{5, 2, 3, 1}
	quickSort(input)
	fmt.Println(input)
}
