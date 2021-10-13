package main

import "fmt"

func sqrt(a int, precision int) float64 {

	left := float64(0)
	right := float64(a)
	//i := 0
	for ; right-left > 0.01; {
		mid := (right + left) / 2

		if mid*mid > float64(a) {
			right = mid
			continue
		}
		left = mid
	}
	fmt.Println(right, left)
	return float64(left)
}

func main() {

	a := 10
	fmt.Println(sqrt(a, 10))

}
