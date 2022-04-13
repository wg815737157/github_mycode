package main

import (
	"fmt"
	"strconv"
	"sync"
)

func testFn3(n int) int {
	res := 0
	for i := n; i >= 1; i-- {
		res += i
	}
	return res
}

func testFn(n int) int {
	if n == 1 {
		return 1
	}
	return testFn(n-1) + n
}

func testFn2(n int) int {
	if n == 1 {
		return 1
	}
	return n * testFn2(n-1)
}
func climbStairs2(n int) int {
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

func climbStairs(n int) int {
	res := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		if i == 1 {
			res[i] = 1
		} else if i == 2 {
			res[i] = 2
		} else {
			res[i] = res[i-1] + res[i-2]
		}
	}
	return res[n]
}

func test() {
	i := 0
	i++
	test()
}

var c chan int
var wg sync.WaitGroup

type TestStruct struct {
	A      int
	B      int
	BddDEE int
}

func main() {
	tmp, err := strconv.ParseInt("000", 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tmp)
	return

	//time.Sleep(10 * time.Second)

	//m := new(map[string]string)
	//
	//fmt.Printf("%T\n", m)
	//
	//var a map[string]string
	//fmt.Printf("%T\n", a)
	//
	//c := make(chan int)
	//fmt.Printf("%T\n", c)
	//
	//i := 0
	//
	//p := &i
	//fmt.Printf("%T\n", p)
	//_ = 1 / 0

	//fmt.Println(testFn(10))
	//fmt.Println(testFn3(10))
}
