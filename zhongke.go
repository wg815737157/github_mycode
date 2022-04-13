package main

import (
	"fmt"
	"sort"
	"time"
)

func solution(a, b []int) []int {
	i, j, k := 0, 0, 0
	m, n := len(a), len(b)
	res := make([]int, m+n)
	for ; i < m && j < n; k++ {
		if a[i] < b[j] {
			res[k] = a[i]
			i++
		} else {
			res[k] = b[j]
			j++
		}
	}
	for i < m {
		res[k] = a[i]
		i++
		k++
	}
	for j < n {
		res[k] = b[j]
		j++
		k++
	}
	return res
}

var mq *Mq

type Mq struct {
	Buffer []int
	ChanA  chan int
}

func producer() {
	for {
		time.Sleep(time.Second)
		mq.ChanA <- 1
	}
}
func consumer() {
	for v := range mq.ChanA {
		fmt.Println(v)
	}
}

func main() {


	mq = &Mq{Buffer: []int{}, ChanA: make(chan int)}
	go consumer()
	go producer()
	select {}

}
