package main

import (
	"fmt"
)

func cancelBlank(s string) {
	runeArray := []rune(s)
	j := 0
	i := 0
	for ; i < len(runeArray)-1; i++ {
		if runeArray[i] == ' ' && runeArray[i] == runeArray[i+1] {
			continue
		}
		runeArray[j] = runeArray[i]
		j++
	}
	//加上结尾
	runeArray[j] = runeArray[i]
	fmt.Println(string(runeArray[:j+1]), "end")
}

func initArrays() [][]int {
	item1 := []int{1, 2, 3, 4, 5}
	item2 := []int{6, 7}
	item3 := []int{8, 9, 10}
	item4 := []int{11}
	res := [][]int{item1, item2, item3, item4}
	return res
}

// 探探 test
func solution(initArray [][]int) {
	row := len(initArray)
	maxColumn := 0
	//获取最大列值
	for i := 0; i < row; i++ {
		if len(initArray[i]) > maxColumn {
			maxColumn = len(initArray[i])
		}
	}

	res := make([][]int, 0, 0)
	for column := 0; column < maxColumn; column++ {
		item := make([]int, 0, 0)
		for i := 0; i < row; i++ {
			iColumn := len(initArray[i])
			if column < iColumn {
				item = append(item, initArray[i][column])
			}
		}
		res = append(res, item)
	}
	fmt.Println(res)
}

//蛇形打印
func initSnakeMatrix() [][]int {
	item1 := []int{1, 2, 3, 4}
	item2 := []int{5, 6, 7, 8}
	item3 := []int{9, 10, 11, 12}
	res := [][]int{item1, item2, item3}
	return res
}

func printSnakeMatrix(input [][]int) {

	left, top := 0, 0
	right := len(input[0])
	bottom := len(input)

	for left < right && top < bottom {
		leftRight := left
		for ; leftRight < right; leftRight++ {
			fmt.Println(input[top][leftRight])
		}
		topBottom := top + 1
		for ; topBottom < bottom; topBottom++ {
			fmt.Println(input[topBottom][right-1])
		}

		rightLeft := right - 2
		for ; rightLeft >= left; rightLeft-- {
			fmt.Println(input[bottom-1][rightLeft])
		}

		bottomTop := bottom - 2
		for ; bottomTop > top; bottomTop-- {
			fmt.Println(input[bottomTop][left])
		}
		right--
		left++
		top++
		bottom--

	}

}

func test(a ...int) {

}

func main() {

	cancelBlank("a b ccc ddd                  eee  ee ")
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("Recovered in f", r)
	//
	//		debug.PrintStack()
	//	}
	//}()
	//
	//a := make([]int, 0, 0)
	//fmt.Println(a[1])

	//initArray := initSnakeMatrix()
	//
	//printSnakeMatrix(initArray)

}
