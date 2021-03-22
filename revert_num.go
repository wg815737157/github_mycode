package main

import "fmt"

func main() {
	var input1 int
	i, err := fmt.Scan(&input1)

	fmt.Println("输入的值为：", input1)
	if err != nil {
		fmt.Println(err)
		return
	}
	if i != 1 {
		fmt.Println(i, ":argument err")
		return
	}
	ret := 0
	x := input1

	for x != 0 {
		pop := x % 10
		x /= 10
		ret = ret*10 + pop
	}
	fmt.Println(ret)
}

