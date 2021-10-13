package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	for i, j := 0, n-1; i < j; {
		if numbers[i]+numbers[j] == target {
			return []int{i+1, j+1}
		}
		if numbers[i]+numbers[j] > target {
			j--
		} else {
			i++
		}
	}
	return nil

}

func main() {

	a := []int{1, 2, 3, 4, 5, 6, 7}

	array := [4]int{1, 2, 4}
	fmt.Println(array)
	fmt.Println(twoSum(a, 10))

}
