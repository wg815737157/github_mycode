package main

import (
	"fmt"
	"math"
	"strconv"
	"sync"
)

func missingTwo(nums []int) []int {
	max := math.MinInt32
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	bitMap := make([]byte, max/8+1)
	for i := 0; i < len(nums); i++ {
		yu := nums[i] % 8
		shang := nums[i] / 8
		bitMap[shang] = bitMap[shang] | (1 << yu)
	}

	res := []int{}

	for i := 0; i < len(bitMap); i++ {
		for n := 0; n < 8; n++ {
			flag := byte(1 << n)
			if bitMap[i]&flag != flag {
				v := i*8 + n
				if v > 0 && v < max {
					res = append(res, v)
				}
			}
		}
	}
	if len(res) == 0 {
		res = append(res, max+1)
		res = append(res, max+2)
	} else if len(res) == 1 {
		res = append(res, max+1)
	}
	return res
}

type der interface {
	execute()
}
type std struct {
	chanM chan der
	i     int
}

func missingTwoII(nums []int) []int {
	return nil
}
func main() {
	//chanI := make(chan int, 10)
	resSlice := []string{}
	wg := &sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			resSlice = append(resSlice, strconv.Itoa(i))
		}(i)
	}
	wg.Wait()
	fmt.Println(resSlice)
	return
}
