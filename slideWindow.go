package main

import (
	"fmt"
	"math"
)



func minSubArrayLen(inputArray []int, target int) int {
	if len(inputArray) == 0 {
		return -1
	}

	minLen := math.MaxInt32
	sum := 0
	left, right := 0, 0
	for ; right < len(inputArray); right++ {
		sum += inputArray[right]
		if sum < target {
			continue
		}
		minLen = minInt(minLen, right-left+1)
		for ; left <= right; left++ {
			sum -= inputArray[left]
			if sum >= target {
				continue
			}
			fmt.Println("left", left)
			fmt.Println("right", right)

			minLen = minInt(minLen, right-left+1)
			left++
			break
		}
	}

	if minLen == math.MaxInt32 {
		return 0
	}
	fmt.Println(minLen)
	return minLen
}

func trans(input, inputN, outputM int) {
	jinzhi10 := 0
	//获取10进制数
	for i := 0; input > 0; i++ {
		yu := input % 10
		input = input / 10
		jinzhi10 += yu * int(math.Pow(float64(inputN), float64(i)))
	}
	fmt.Println(jinzhi10)
	s := &Stack{}
	for jinzhi10 > 0 {
		s.Push(jinzhi10 % 2)
		jinzhi10 /= 2
	}
	for s.Len() > 0 {
		fmt.Println(s.Pop())
	}
}
func main() {
	trans(721, 8, 2)
}
