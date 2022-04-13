package main

import (
	"fmt"
)

func main() {


	prices := []int{1, 3, 2, 4, 9, 3}
	fmt.Println(maxProfile(prices))
	fmt.Println(maxProfile2(prices))
	fmt.Println(maxProfile3(prices))

}
