package main

import (
	"fmt"
)

func main() {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		fmt.Println(err)
		return
	}
	intput := make([]int64, n)
	for i := 0; i < n; i++ {
		var v int64
		_, err := fmt.Scanln(&v)
		if err != nil {
			fmt.Println(err)
			return
		}
		intput[i] = v
	}
	output := solution(intput)
	for i := 0; i < len(output); i++ {
		fmt.Println(output[i])
	}
}

//
func solution(input []int64) []int {
	n := len(input)
	if n == 0 {
		return nil
	}
	res := []int{}
	for i := 0; i < n; i++ {
		digits := 1
		input[i] /= 10
		for input[i] != 0 {
			digits++
			input[i] /= 10
		}
		res = append(res, digits)
	}
	return res
}
